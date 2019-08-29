# Error codes for Bleu API

This is the error JSON:

```javascript

{

  "code": "1102",

  "message": "Invalid 2FA code.",

  "HTTP code": "401"

}

```

## 10xx - Server or Network error

#### 1000 UNKNOWN

 * Unknown request error.
 * HTTP code: 502

#### 1001 DISCONNECTED

 * Internal error. Please try again.
 * HTTP code: 500

## 11xx - Request issues

#### 1100 UNKNOWN PARAMETERS

 * An unknown parameter was sent.
 * HTTP code: 422

#### 1101 INVALID COUNTRY

 * Invalid country.
 * HTTP code: 403

#### 1102 INVALID 2FA

 * Invalid 2FA code.
 * HTTP code: 401

#### 1103 INVALID E-MAIL CODE

 * Invalid e-mail 2FA code.
 * HTTP code: 401

#### 1104 INVALID CPF

 * CPF already exists.
 * HTTP code: 409

#### 1105 INVALID USERNAME

 * Username already exists.
 * HTTP code: 409

#### 1106 INVALID LOGIN

 * Invalid Login.
 * HTTP code: 401

#### 1109 INVALID QUANTITY OF ATTEMPTS

 * Too many attempts.
 * HTTP code: 429
