package com.fluxninja.aperture.servlet.jakarta;

import com.fluxninja.aperture.sdk.ApertureSDKException;
import com.fluxninja.aperture.sdk.FlowStatus;
import com.fluxninja.aperture.sdk.TrafficFlow;
import com.fluxninja.aperture.sdk.Utils;
import com.fluxninja.generated.aperture.flowcontrol.checkhttp.v1.CheckHTTPRequest;
import io.opentelemetry.api.baggage.Baggage;
import io.opentelemetry.api.baggage.BaggageEntry;
import jakarta.servlet.ServletRequest;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletRequestWrapper;
import java.net.URLDecoder;
import java.nio.charset.StandardCharsets;
import java.util.*;

public class ServletUtils {
    protected static int handleRejectedFlow(TrafficFlow flow) {
        try {
            flow.end(FlowStatus.Unset);
        } catch (ApertureSDKException e) {
            e.printStackTrace();
        }
        if (flow.checkResponse() != null
                && flow.checkResponse().hasDeniedResponse()
                && flow.checkResponse().getDeniedResponse().getStatus() != 0) {
            return flow.checkResponse().getDeniedResponse().getStatus();
        }
        return 403;
    }

    protected static CheckHTTPRequest checkRequestFromRequest(ServletRequest req) {
        Map<String, String> baggageLabels = new HashMap<>();

        for (Map.Entry<String, BaggageEntry> entry : Baggage.current().asMap().entrySet()) {
            String value;
            try {
                value =
                        URLDecoder.decode(
                                entry.getValue().getValue(), StandardCharsets.UTF_8.name());
            } catch (java.io.UnsupportedEncodingException e) {
                // This should never happen, as `StandardCharsets.UTF_8.name()` is a valid
                // encoding
                throw new RuntimeException(e);
            }
            baggageLabels.put(entry.getKey(), value);
        }

        return addHttpAttributes(baggageLabels, req).build();
    }

    protected static ServletRequest updateHeaders(
            ServletRequest req, Map<String, String> newHeaders) {
        HttpServletRequest httpReq = (HttpServletRequest) req;
        Map<String, String> headerMap = new HashMap<>(newHeaders);
        return new HttpServletRequestWrapper(httpReq) {
            @Override
            public Enumeration<String> getHeaderNames() {
                Set<String> headerNames = new HashSet<>(Collections.list(super.getHeaderNames()));
                headerNames.addAll(headerMap.keySet());
                return Collections.enumeration(headerNames);
            }

            @Override
            public String getHeader(String name) {
                String header = headerMap.get(name);
                return header != null ? header : super.getHeader(name);
            }

            @Override
            public Enumeration<String> getHeaders(String name) {
                String header = headerMap.get(name);
                if (header != null) {
                    List<String> values = Arrays.asList(header.split(","));
                    return Collections.enumeration(values);
                } else {
                    return super.getHeaders(name);
                }
            }
        };
    }

    private static CheckHTTPRequest.Builder addHttpAttributes(
            Map<String, String> headers, ServletRequest req) {
        HttpServletRequest request = (HttpServletRequest) req;
        Enumeration<String> originalHeaders = request.getHeaderNames();
        while (originalHeaders.hasMoreElements()) {
            String headerKey = originalHeaders.nextElement();
            headers.put(headerKey, request.getHeader(headerKey));
        }

        String sourceIp = req.getRemoteAddr();
        int sourcePort = req.getRemotePort();
        String destinationIp = req.getLocalAddr();
        int destinationPort = req.getLocalPort();

        CheckHTTPRequest.Builder builder = CheckHTTPRequest.newBuilder();

        builder.setControlPoint("ingress")
                .setRequest(
                        CheckHTTPRequest.HttpRequest.newBuilder()
                                .setMethod(request.getMethod())
                                .setPath(request.getServletPath())
                                .setHost(req.getRemoteHost())
                                .setScheme(req.getScheme())
                                .setSize(req.getContentLength())
                                .setProtocol(req.getProtocol())
                                .putAllHeaders(headers));

        if (sourceIp != null) {
            builder.setSource(Utils.createSocketAddress(sourceIp, sourcePort, "TCP"));
        }
        if (destinationIp != null) {
            builder.setDestination(
                    Utils.createSocketAddress(destinationIp, destinationPort, "TCP"));
        }
        return builder;
    }
}
