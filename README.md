# Error codes for Bleu API

This is the error JSON:

```javascript

{

  "code": "1102",

  "Error Type": "request-issue",

  "HTTP code": "401",

  "message": "Invalid 2FA code."

}

```

## 10xx - Server or Network error

#### 1000 UNKNOWN

 * Unknown request error.
 * HTTP code: 502
 * Error Type: system-issue

#### 1001 DISCONNECTED

 * Internal error. Please try again.
 * HTTP code: 500
 * Error Type: system-issue

## 11xx - Request issues

#### 1100 UNKNOWN PARAMETERS

 * An unknown parameter was sent.
 * HTTP code: 422
 * Error Type: request-issue

#### 1101 INVALID COUNTRY

 * Invalid country.
 * HTTP code: 403
 * Error Type: request-issue

#### 1102 INVALID 2FA

 * Invalid 2FA code.
 * HTTP code: 401
 * Error Type: request-issue

#### 1103 INVALID E-MAIL CODE

 * Invalid e-mail 2FA code.
 * HTTP code: 401
 * Error Type: request-issue

#### 1104 INVALID CPF

 * CPF already exists.
 * HTTP code: 409
 * Error Type: request-issue

#### 1105 INVALID USERNAME

 * Username already exists.
 * HTTP code: 409
 * Error Type: request-issue

#### 1106 INVALID LOGIN

 * Invalid Login.
 * HTTP code: 401
 * Error Type: request-issue

#### 1109 INVALID QUANTITY OF ATTEMPTS

 * Too many attempts.
 * HTTP code: 429
 * Error Type: request-issue
