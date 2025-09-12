# API Logging Standards and Guidelines

## Overview

This document establishes the standard logging practices for all API development projects. Consistent logging is crucial for monitoring, debugging, troubleshooting, and maintaining production systems.

## Mandatory Logging Rule

**RULE: All HTTP APIs MUST implement request/response logging using the standardized Apache-style format with response time tracking.**

## Standard Log Format

### Format Specification
```
{host} - - {timestamp} "{method} {path} {protocol}" {status_code} {content_length} "-" "{user_agent}" {response_time_ms}
```

### Implementation Template
```go
log.Printf("%s - - %s \"%s %s %s\" %d %d \"-\" \"%s\" %d\n",
    r.Host,                                    // Client host/IP address
    t.Format("[02/Jan/2006:15:04:05 -0700]"), // Timestamp in Apache format
    r.Method,                                  // HTTP method (GET, POST, PUT, DELETE, etc.)
    r.URL.Path,                               // Request path
    r.Proto,                                  // HTTP protocol version
    statusCode,                               // HTTP response status code
    r.ContentLength,                          // Request content length in bytes
    r.UserAgent(),                           // Client user agent string
    time.Since(startTime).Milliseconds(),    // Response time in milliseconds
)
```

### Example Output
```
api.example.com - - [15/Jan/2024:14:30:25 +0000] "GET /api/v1/users HTTP/1.1" 200 0 "-" "curl/7.68.0" 15
api.example.com - - [15/Jan/2024:14:30:30 +0000] "POST /api/v1/users HTTP/1.1" 201 85 "-" "PostmanRuntime/7.28.4" 142
api.example.com - - [15/Jan/2024:14:30:35 +0000] "GET /api/v1/users/123 HTTP/1.1" 404 45 "-" "Mozilla/5.0" 8
```

## Implementation Requirements

### 1. Middleware Implementation
- **MUST** be implemented as HTTP middleware
- **MUST** wrap the response writer to capture status codes
- **MUST** record start time before processing request
- **MUST** log after request completion

### 2. Required Data Points
All log entries **MUST** include:
- Client host/IP address
- Timestamp in Apache format (`[02/Jan/2006:15:04:05 -0700]`)
- HTTP method
- Request path (without query parameters for security)
- HTTP protocol version
- Response status code
- Request content length
- User agent string
- Response time in milliseconds

### 3. Timing Precision
- Response time **MUST** be measured in milliseconds
- Timing **MUST** start before request processing begins
- Timing **MUST** end after response is fully written

## Code Implementation Standards

### Go Implementation Example
```go
// responseWriter wraps http.ResponseWriter to capture status code
type responseWriter struct {
    http.ResponseWriter
    statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
    rw.statusCode = code
    rw.ResponseWriter.WriteHeader(code)
}

// loggingMiddleware implements standard API logging
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        startTime := time.Now()
        
        // Wrap response writer to capture status code
        wrapped := &responseWriter{ResponseWriter: w, statusCode: 200}
        
        // Process request
        next.ServeHTTP(wrapped, r)
        
        // Log in standard format
        log.Printf("%s - - %s \"%s %s %s\" %d %d \"-\" \"%s\" %d\n",
            r.Host,
            startTime.Format("[02/Jan/2006:15:04:05 -0700]"),
            r.Method,
            r.URL.Path,
            r.Proto,
            wrapped.statusCode,
            r.ContentLength,
            r.UserAgent(),
            time.Since(startTime).Milliseconds(),
        )
    })
}
```

## Security Considerations

### What to Log
- Host/IP addresses (for rate limiting and security monitoring)
- HTTP methods and paths
- Status codes
- Response times
- User agents (for bot detection)

### What NOT to Log
- **NEVER** log sensitive data in URLs (passwords, tokens, API keys)
- **NEVER** log request/response bodies containing PII
- **NEVER** log authentication headers or cookies
- **NEVER** log query parameters that might contain sensitive data

### Query Parameter Handling
- Use `r.URL.Path` instead of `r.URL.String()` to avoid logging query parameters
- If query parameter logging is required, implement allowlist filtering

## Performance Requirements

### Response Time Thresholds
- **< 100ms**: Excellent performance
- **100-500ms**: Good performance
- **500-1000ms**: Acceptable performance (monitor closely)
- **> 1000ms**: Poor performance (requires investigation)

### Log Volume Management
- Implement log rotation for high-traffic APIs
- Consider sampling for extremely high-volume endpoints (> 1000 req/sec)
- Use appropriate log levels (INFO for standard requests, WARN for 4xx, ERROR for 5xx)

## Monitoring and Alerting

### Metrics to Extract
From the standard log format, extract:
- Request rate (requests per second/minute/hour)
- Error rate (4xx/5xx responses)
- Average/P95/P99 response times
- Most common endpoints
- User agent patterns (for bot detection)

### Alert Conditions
Set up alerts for:
- Error rate > 5% over 5 minutes
- P95 response time > 1000ms over 5 minutes
- Request rate anomalies (sudden spikes/drops)
- High number of 401/403 responses (potential security issues)

## Framework-Specific Guidelines

### Go (Gorilla Mux, Chi, Gin)
- Implement as middleware function
- Register middleware before route handlers
- Use `http.ResponseWriter` wrapper pattern

### Node.js (Express, Fastify, Koa)
- Implement as middleware function
- Use response interceptor or event listeners
- Handle async/await patterns properly

### Python (Flask, Django, FastAPI)
- Implement as decorator or middleware class
- Use response hooks or middleware pipeline
- Handle both sync and async endpoints

### Java (Spring Boot, Jersey)
- Implement as Filter or Interceptor
- Use HandlerInterceptor for Spring Boot
- Implement ContainerRequestFilter for Jersey

## Compliance and Auditing

### Regulatory Requirements
- Ensure logging complies with GDPR, CCPA, and other privacy regulations
- Implement data retention policies
- Provide audit trails for compliance reporting

### Log Retention
- **Development**: 7 days minimum
- **Staging**: 30 days minimum
- **Production**: 90 days minimum (adjust based on compliance requirements)

## Exceptions and Variations

### When Standard Format May Be Modified
1. **Healthcare APIs**: May require additional HIPAA compliance logging
2. **Financial APIs**: May require additional audit fields for PCI compliance
3. **High-frequency APIs**: May require sampling or reduced logging

### Approval Process
Any deviation from this standard **MUST**:
1. Be documented with business justification
2. Be approved by the architecture review board
3. Maintain equivalent monitoring capabilities
4. Include migration plan to standard format

## Tools and Libraries

### Recommended Logging Libraries
- **Go**: `log/slog`, `logrus`, `zap`
- **Node.js**: `winston`, `pino`
- **Python**: `logging`, `structlog`
- **Java**: `logback`, `log4j2`

### Log Aggregation
- Use centralized logging (ELK Stack, Splunk, CloudWatch)
- Implement structured logging where possible
- Tag logs with service name and version

## Review and Updates

This document should be:
- Reviewed quarterly by the architecture team
- Updated when new requirements emerge
- Versioned with changelog tracking

## Enforcement

### Code Review Requirements
- All API implementations **MUST** include logging middleware
- Code reviews **MUST** verify standard log format compliance
- Automated tests **SHOULD** verify log format output

### Monitoring Compliance
- Automated checks for missing request logging
- Regular audits of log format consistency
- Performance monitoring of logging overhead

---

**Document Version**: 1.0  
**Last Updated**: January 2024  
**Next Review**: April 2024  
**Owner**: Platform Architecture Team