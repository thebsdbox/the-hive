package k3d

var cilium = `# Source: cilium/templates/cilium-agent/serviceaccount.yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: "cilium"
  namespace: kube-system
---
# Source: cilium/templates/cilium-operator/serviceaccount.yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: "cilium-operator"
  namespace: kube-system
---
# Source: cilium/templates/hubble-relay/serviceaccount.yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: "hubble-relay"
  namespace: kube-system
---
# Source: cilium/templates/hubble-ui/serviceaccount.yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: "hubble-ui"
  namespace: kube-system
---
# Source: cilium/templates/cilium-ca-secret.yaml
apiVersion: v1
kind: Secret
metadata:
  name: cilium-ca
  namespace: kube-system
data:
  ca.crt: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURGRENDQWZ5Z0F3SUJBZ0lSQUxIeHRLQ2RWaUR3cU1JS1JMY1NCR0V3RFFZSktvWklodmNOQVFFTEJRQXcKRkRFU01CQUdBMVVFQXhNSlEybHNhWFZ0SUVOQk1CNFhEVEl6TURNeE56QTVNVE0xTTFvWERUSTJNRE14TmpBNQpNVE0xTTFvd0ZERVNNQkFHQTFVRUF4TUpRMmxzYVhWdElFTkJNSUlCSWpBTkJna3Foa2lHOXcwQkFRRUZBQU9DCkFROEFNSUlCQ2dLQ0FRRUF3VFJFeXJldnlsMm9XSFZaYnkvMXBmUXFscXVTVTlyV0dUMUxOd3NNcE96UlFIZi8KejlLRHZrTnVGSXZDZ2tiS0RPY1J2UkVvTFlsZGpjWGZwTFRrbFlUYnBjRFJReHd3elRXK3hlWnZSV2tndVpQUQo3WENZV3RrZ2FXcUQ5bjduMFRSU3FMdTU5TEo4L3BoYlYvYlVHSGZ0bGxEUVZCbXZ5N0Y2S3QzdDJRamw2aFNNCmY0VUhwZjFpMEdWb21qUzZWeDZvMmQxeFhNUEh6ZGVoYjQ4Wk1QYzA5UGQySUhCYlRsMk1wZjF6Z05Pb3l6MnYKZUwrSlIzV0FFY3lySVNBRU02bGtnRzZ6blEzSnV2Z0xwcHBtZ1pMbEYzVXlDYXgxTVNGMGkxb1RLNEtIWDQyaApsT3RkUyt2REl1VWVUeFRjUWxyRmdMT0hBR3RLMDQ2aTQrMVZuUUlEQVFBQm8yRXdYekFPQmdOVkhROEJBZjhFCkJBTUNBcVF3SFFZRFZSMGxCQll3RkFZSUt3WUJCUVVIQXdFR0NDc0dBUVVGQndNQ01BOEdBMVVkRXdFQi93UUYKTUFNQkFmOHdIUVlEVlIwT0JCWUVGTG5rNDEwR0Z0Qy96TUtkN2xLa2RCYm4rcE9tTUEwR0NTcUdTSWIzRFFFQgpDd1VBQTRJQkFRQW5KZHhRZXliRFpxeGQxYllMazZaeDRkYVo2Tkc3eDlyZ2lUMWgyT21tOFBFUE5OZngwa090CnZ0L0xpSGltUTlxMDl2c0pleERLQU1vL1h6Z3JxZ0pvbkNrY1RVTk1vdlVrU2loM2IxSVdSYXFNNUtRemxsN1QKbnRwem5SQXZyTElGamNUYlpIY2JMYk5qVkprY2xwWFdTYTFkVnlMOElKbVFoaUljSnQyS041TFNPdnRRbFdKTwprWFNtdVEwbUFoTW0wVEVyWnM3dWRRQ2I5RDFrd0svQlQwVU84Q2dheEZwZEZET2g0TGJ5NERMN0dnc25RY1ZGCld4UHl6elExdHdDRkxPa2dxYnVhZ25TWUdYVW9XbjRNMzREVE1USUk0N1ZNVDgzYzR6Q2hoU1RTanNLRDFxaCsKQWthUTlOazE1L01QUjExRkhVKytTeld4eGdEK0hkanoKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=
  ca.key: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcFFJQkFBS0NBUUVBd1RSRXlyZXZ5bDJvV0hWWmJ5LzFwZlFxbHF1U1U5cldHVDFMTndzTXBPelJRSGYvCno5S0R2a051Rkl2Q2drYktET2NSdlJFb0xZbGRqY1hmcExUa2xZVGJwY0RSUXh3d3pUVyt4ZVp2UldrZ3VaUFEKN1hDWVd0a2dhV3FEOW43bjBUUlNxTHU1OUxKOC9waGJWL2JVR0hmdGxsRFFWQm12eTdGNkt0M3QyUWpsNmhTTQpmNFVIcGYxaTBHVm9talM2Vng2bzJkMXhYTVBIemRlaGI0OFpNUGMwOVBkMklIQmJUbDJNcGYxemdOT295ejJ2CmVMK0pSM1dBRWN5cklTQUVNNmxrZ0c2em5RM0p1dmdMcHBwbWdaTGxGM1V5Q2F4MU1TRjBpMW9USzRLSFg0MmgKbE90ZFMrdkRJdVVlVHhUY1FsckZnTE9IQUd0SzA0Nmk0KzFWblFJREFRQUJBb0lCQVFDbWErT2x5cnY5ME5jbAorMDlpV3lsNkhleGhoTUN5d3BvUmU1N0NkeDRaSW5UTTE1RGU3ekxVUFRraE1aSUdyZnQxU2NVWjJ5R0V1QUtDCkVLUW9PRGwrSWNjaDdBbFZpWDd3QlErallHOFFseEc5U1BwNnpMWXY2ei9rRGdCRjMrTEhiZzZWa2I4SWJJY0cKdHk1WGd5SUhRYWJDQ0RLcmZzVFM3TG1UR3orNTVsUVZ2VUVpSlJ5MWxPeGdYRm9jZEE1aUZDUFZCRTU4ZUJqOQpsOURSaFNUVFZrVVlWUG1HdEY0SHB6Z2NJU25BSFpadzFLY1VhcXJzSVkwamZMZ3c1YXNxNXlNUE50VFFubEFwCnc3WktQR1QzVEdPZExlRWpQclQ0bGhYY0p1eHZpMThubTJyZHA0N3l2V3AxZ0JaQVRwYVJkN01Nd3AxejV5ZHcKNTFuYVFwR0JBb0dCQU02UEpXcG9uSmJSMkZIb1M5aDMzQ3kyVTd2R2dlL2hTdm1rNzdYS2NEUllCU3dsWGs5dApndFhndTZ1TUtOUWRaeEZBd1BGWlVqa09kcTB3OUoramRTY09EVThJSWh3NzdudFJLbVRQVlliaE5qY013UlQ3CnowQkd5ZlZkZnM3T0NqVVVydXArdXMxUVByV2NBcnVCYVoyNGRhcGFpUWoxb3BuNXNTdjRlZnZaQW9HQkFPOXkKektNblFtcTR2ejNDWW90Y0lMNXVhSXhIYjB4TFd0QjkrY0kxY2p2QUs5ZW81ZEZNOEtmdDJubVoyaHc4UTdIQwphalJuZ1RKeGYwZlJVNDZPNUM4TEZGLzFGc28rbHNsT2hicW9oNjhvMnFtSlFUVm8veVQwdldrVFIyU05HczBWCkFMSndLdlJrVWVwZ3VjODRZTzdEQVYxZnBGTHpLZkJJcnY5TlJpRmxBb0dBQjYyQ2NvWVk2L0k0M0RLS1B5MlYKWFlRWmNLMWNQeEpjdXhMS1pqTjBJRDMxVTBMQVVxdDdaWC9JK2dObnNScTJyZ2wrSW5wemQvTjFyZEpZQldjSgovNzJoK1FJUVlvUkh4UVdyVWJ2ekxlUkpJNXF4d3BucGhqWWJZNmRxQXozZFcwTzlqTEhSTjdoMzNFQkVTYnZ4CnRROGFNSTdVOFNSUU92RHhDUFZmYzJFQ2dZRUE2UktsZm1wSWkva29yY1Q0aHc0MkVTY0hQUVNMb1lmMzdkbXgKc3dpekdOWUYxdlhnUGNyV3RaOGdlaHozNFdRSHdJK3RNVFZPM1ByOUdicjN5bHZzWUo0NFJ1OGFMK0tjZzNhYgpWUVdXalRrSEh0OHJTZ0haMk84aEw1WkVkK3VobXQ1R3YyblBaZlFBaUZOK2llWW05RUY4b3BibUxKZmt5cTcxCktDemZoc0VDZ1lFQXR4RXMvTkN0Z2dFL1hOMG1xOFB3eFp2SUt6bjBrVHYwWi9NYnoxUm1IT1NzaCtiT3QxTloKM3VXY3dmbGlVSmhqV3cyaEdSYjROU3VTUGVWTzdNTWxPUlpTZXFwZnc3Z3BKNTBDZ2lSTktNUStDUDhUUHdWUQppdjVhRlMzRXg0a1o5bmU0K25oV0kwczFwWlBkYlRpdzVVTzJmRmMzaEE0TkpqWGlJdE1jK3hRPQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo=
---
# Source: cilium/templates/hubble/tls-helm/ca-secret.yaml
apiVersion: v1
kind: Secret
metadata:
  name: hubble-ca-secret
  namespace: kube-system
data:
  ca.crt: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURGRENDQWZ5Z0F3SUJBZ0lSQUxIeHRLQ2RWaUR3cU1JS1JMY1NCR0V3RFFZSktvWklodmNOQVFFTEJRQXcKRkRFU01CQUdBMVVFQXhNSlEybHNhWFZ0SUVOQk1CNFhEVEl6TURNeE56QTVNVE0xTTFvWERUSTJNRE14TmpBNQpNVE0xTTFvd0ZERVNNQkFHQTFVRUF4TUpRMmxzYVhWdElFTkJNSUlCSWpBTkJna3Foa2lHOXcwQkFRRUZBQU9DCkFROEFNSUlCQ2dLQ0FRRUF3VFJFeXJldnlsMm9XSFZaYnkvMXBmUXFscXVTVTlyV0dUMUxOd3NNcE96UlFIZi8KejlLRHZrTnVGSXZDZ2tiS0RPY1J2UkVvTFlsZGpjWGZwTFRrbFlUYnBjRFJReHd3elRXK3hlWnZSV2tndVpQUQo3WENZV3RrZ2FXcUQ5bjduMFRSU3FMdTU5TEo4L3BoYlYvYlVHSGZ0bGxEUVZCbXZ5N0Y2S3QzdDJRamw2aFNNCmY0VUhwZjFpMEdWb21qUzZWeDZvMmQxeFhNUEh6ZGVoYjQ4Wk1QYzA5UGQySUhCYlRsMk1wZjF6Z05Pb3l6MnYKZUwrSlIzV0FFY3lySVNBRU02bGtnRzZ6blEzSnV2Z0xwcHBtZ1pMbEYzVXlDYXgxTVNGMGkxb1RLNEtIWDQyaApsT3RkUyt2REl1VWVUeFRjUWxyRmdMT0hBR3RLMDQ2aTQrMVZuUUlEQVFBQm8yRXdYekFPQmdOVkhROEJBZjhFCkJBTUNBcVF3SFFZRFZSMGxCQll3RkFZSUt3WUJCUVVIQXdFR0NDc0dBUVVGQndNQ01BOEdBMVVkRXdFQi93UUYKTUFNQkFmOHdIUVlEVlIwT0JCWUVGTG5rNDEwR0Z0Qy96TUtkN2xLa2RCYm4rcE9tTUEwR0NTcUdTSWIzRFFFQgpDd1VBQTRJQkFRQW5KZHhRZXliRFpxeGQxYllMazZaeDRkYVo2Tkc3eDlyZ2lUMWgyT21tOFBFUE5OZngwa090CnZ0L0xpSGltUTlxMDl2c0pleERLQU1vL1h6Z3JxZ0pvbkNrY1RVTk1vdlVrU2loM2IxSVdSYXFNNUtRemxsN1QKbnRwem5SQXZyTElGamNUYlpIY2JMYk5qVkprY2xwWFdTYTFkVnlMOElKbVFoaUljSnQyS041TFNPdnRRbFdKTwprWFNtdVEwbUFoTW0wVEVyWnM3dWRRQ2I5RDFrd0svQlQwVU84Q2dheEZwZEZET2g0TGJ5NERMN0dnc25RY1ZGCld4UHl6elExdHdDRkxPa2dxYnVhZ25TWUdYVW9XbjRNMzREVE1USUk0N1ZNVDgzYzR6Q2hoU1RTanNLRDFxaCsKQWthUTlOazE1L01QUjExRkhVKytTeld4eGdEK0hkanoKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=
  ca.key: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcFFJQkFBS0NBUUVBd1RSRXlyZXZ5bDJvV0hWWmJ5LzFwZlFxbHF1U1U5cldHVDFMTndzTXBPelJRSGYvCno5S0R2a051Rkl2Q2drYktET2NSdlJFb0xZbGRqY1hmcExUa2xZVGJwY0RSUXh3d3pUVyt4ZVp2UldrZ3VaUFEKN1hDWVd0a2dhV3FEOW43bjBUUlNxTHU1OUxKOC9waGJWL2JVR0hmdGxsRFFWQm12eTdGNkt0M3QyUWpsNmhTTQpmNFVIcGYxaTBHVm9talM2Vng2bzJkMXhYTVBIemRlaGI0OFpNUGMwOVBkMklIQmJUbDJNcGYxemdOT295ejJ2CmVMK0pSM1dBRWN5cklTQUVNNmxrZ0c2em5RM0p1dmdMcHBwbWdaTGxGM1V5Q2F4MU1TRjBpMW9USzRLSFg0MmgKbE90ZFMrdkRJdVVlVHhUY1FsckZnTE9IQUd0SzA0Nmk0KzFWblFJREFRQUJBb0lCQVFDbWErT2x5cnY5ME5jbAorMDlpV3lsNkhleGhoTUN5d3BvUmU1N0NkeDRaSW5UTTE1RGU3ekxVUFRraE1aSUdyZnQxU2NVWjJ5R0V1QUtDCkVLUW9PRGwrSWNjaDdBbFZpWDd3QlErallHOFFseEc5U1BwNnpMWXY2ei9rRGdCRjMrTEhiZzZWa2I4SWJJY0cKdHk1WGd5SUhRYWJDQ0RLcmZzVFM3TG1UR3orNTVsUVZ2VUVpSlJ5MWxPeGdYRm9jZEE1aUZDUFZCRTU4ZUJqOQpsOURSaFNUVFZrVVlWUG1HdEY0SHB6Z2NJU25BSFpadzFLY1VhcXJzSVkwamZMZ3c1YXNxNXlNUE50VFFubEFwCnc3WktQR1QzVEdPZExlRWpQclQ0bGhYY0p1eHZpMThubTJyZHA0N3l2V3AxZ0JaQVRwYVJkN01Nd3AxejV5ZHcKNTFuYVFwR0JBb0dCQU02UEpXcG9uSmJSMkZIb1M5aDMzQ3kyVTd2R2dlL2hTdm1rNzdYS2NEUllCU3dsWGs5dApndFhndTZ1TUtOUWRaeEZBd1BGWlVqa09kcTB3OUoramRTY09EVThJSWh3NzdudFJLbVRQVlliaE5qY013UlQ3CnowQkd5ZlZkZnM3T0NqVVVydXArdXMxUVByV2NBcnVCYVoyNGRhcGFpUWoxb3BuNXNTdjRlZnZaQW9HQkFPOXkKektNblFtcTR2ejNDWW90Y0lMNXVhSXhIYjB4TFd0QjkrY0kxY2p2QUs5ZW81ZEZNOEtmdDJubVoyaHc4UTdIQwphalJuZ1RKeGYwZlJVNDZPNUM4TEZGLzFGc28rbHNsT2hicW9oNjhvMnFtSlFUVm8veVQwdldrVFIyU05HczBWCkFMSndLdlJrVWVwZ3VjODRZTzdEQVYxZnBGTHpLZkJJcnY5TlJpRmxBb0dBQjYyQ2NvWVk2L0k0M0RLS1B5MlYKWFlRWmNLMWNQeEpjdXhMS1pqTjBJRDMxVTBMQVVxdDdaWC9JK2dObnNScTJyZ2wrSW5wemQvTjFyZEpZQldjSgovNzJoK1FJUVlvUkh4UVdyVWJ2ekxlUkpJNXF4d3BucGhqWWJZNmRxQXozZFcwTzlqTEhSTjdoMzNFQkVTYnZ4CnRROGFNSTdVOFNSUU92RHhDUFZmYzJFQ2dZRUE2UktsZm1wSWkva29yY1Q0aHc0MkVTY0hQUVNMb1lmMzdkbXgKc3dpekdOWUYxdlhnUGNyV3RaOGdlaHozNFdRSHdJK3RNVFZPM1ByOUdicjN5bHZzWUo0NFJ1OGFMK0tjZzNhYgpWUVdXalRrSEh0OHJTZ0haMk84aEw1WkVkK3VobXQ1R3YyblBaZlFBaUZOK2llWW05RUY4b3BibUxKZmt5cTcxCktDemZoc0VDZ1lFQXR4RXMvTkN0Z2dFL1hOMG1xOFB3eFp2SUt6bjBrVHYwWi9NYnoxUm1IT1NzaCtiT3QxTloKM3VXY3dmbGlVSmhqV3cyaEdSYjROU3VTUGVWTzdNTWxPUlpTZXFwZnc3Z3BKNTBDZ2lSTktNUStDUDhUUHdWUQppdjVhRlMzRXg0a1o5bmU0K25oV0kwczFwWlBkYlRpdzVVTzJmRmMzaEE0TkpqWGlJdE1jK3hRPQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo=
---
# Source: cilium/templates/hubble/tls-helm/relay-client-secret.yaml
apiVersion: v1
kind: Secret
metadata:
  name: hubble-relay-client-certs
  namespace: kube-system
type: kubernetes.io/tls
data:
  ca.crt:  LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURGRENDQWZ5Z0F3SUJBZ0lSQUxIeHRLQ2RWaUR3cU1JS1JMY1NCR0V3RFFZSktvWklodmNOQVFFTEJRQXcKRkRFU01CQUdBMVVFQXhNSlEybHNhWFZ0SUVOQk1CNFhEVEl6TURNeE56QTVNVE0xTTFvWERUSTJNRE14TmpBNQpNVE0xTTFvd0ZERVNNQkFHQTFVRUF4TUpRMmxzYVhWdElFTkJNSUlCSWpBTkJna3Foa2lHOXcwQkFRRUZBQU9DCkFROEFNSUlCQ2dLQ0FRRUF3VFJFeXJldnlsMm9XSFZaYnkvMXBmUXFscXVTVTlyV0dUMUxOd3NNcE96UlFIZi8KejlLRHZrTnVGSXZDZ2tiS0RPY1J2UkVvTFlsZGpjWGZwTFRrbFlUYnBjRFJReHd3elRXK3hlWnZSV2tndVpQUQo3WENZV3RrZ2FXcUQ5bjduMFRSU3FMdTU5TEo4L3BoYlYvYlVHSGZ0bGxEUVZCbXZ5N0Y2S3QzdDJRamw2aFNNCmY0VUhwZjFpMEdWb21qUzZWeDZvMmQxeFhNUEh6ZGVoYjQ4Wk1QYzA5UGQySUhCYlRsMk1wZjF6Z05Pb3l6MnYKZUwrSlIzV0FFY3lySVNBRU02bGtnRzZ6blEzSnV2Z0xwcHBtZ1pMbEYzVXlDYXgxTVNGMGkxb1RLNEtIWDQyaApsT3RkUyt2REl1VWVUeFRjUWxyRmdMT0hBR3RLMDQ2aTQrMVZuUUlEQVFBQm8yRXdYekFPQmdOVkhROEJBZjhFCkJBTUNBcVF3SFFZRFZSMGxCQll3RkFZSUt3WUJCUVVIQXdFR0NDc0dBUVVGQndNQ01BOEdBMVVkRXdFQi93UUYKTUFNQkFmOHdIUVlEVlIwT0JCWUVGTG5rNDEwR0Z0Qy96TUtkN2xLa2RCYm4rcE9tTUEwR0NTcUdTSWIzRFFFQgpDd1VBQTRJQkFRQW5KZHhRZXliRFpxeGQxYllMazZaeDRkYVo2Tkc3eDlyZ2lUMWgyT21tOFBFUE5OZngwa090CnZ0L0xpSGltUTlxMDl2c0pleERLQU1vL1h6Z3JxZ0pvbkNrY1RVTk1vdlVrU2loM2IxSVdSYXFNNUtRemxsN1QKbnRwem5SQXZyTElGamNUYlpIY2JMYk5qVkprY2xwWFdTYTFkVnlMOElKbVFoaUljSnQyS041TFNPdnRRbFdKTwprWFNtdVEwbUFoTW0wVEVyWnM3dWRRQ2I5RDFrd0svQlQwVU84Q2dheEZwZEZET2g0TGJ5NERMN0dnc25RY1ZGCld4UHl6elExdHdDRkxPa2dxYnVhZ25TWUdYVW9XbjRNMzREVE1USUk0N1ZNVDgzYzR6Q2hoU1RTanNLRDFxaCsKQWthUTlOazE1L01QUjExRkhVKytTeld4eGdEK0hkanoKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=
  tls.crt: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURTRENDQWpDZ0F3SUJBZ0lRU3VJNHZGbjNTQk1QYkFSUm9DQU9zekFOQmdrcWhraUc5dzBCQVFzRkFEQVUKTVJJd0VBWURWUVFERXdsRGFXeHBkVzBnUTBFd0hoY05Nak13TXpFM01Ea3hNelV6V2hjTk1qWXdNekUyTURreApNelV6V2pBak1TRXdId1lEVlFRRERCZ3FMbWgxWW1Kc1pTMXlaV3hoZVM1amFXeHBkVzB1YVc4d2dnRWlNQTBHCkNTcUdTSWIzRFFFQkFRVUFBNElCRHdBd2dnRUtBb0lCQVFERUMwWDVBZ3ZlRFJycSt1eDFLQXB4bFRqem5uV2MKRHZuTmtPYUdpMlBMSExPM0tYcjJxZzdKZDZSVzRSQXlZTnFVUzFWMkJOTmhhTGZMWVIyWUpaeG5FbkxUR3BCRgp0OUpDbDVESUdkSnpHMm50M0pRV1kwQjM2NkZLR0VqS1pnZUt3U3NmUEUyMWg2dWxZK2hKU0J6OVdRcko4RFRPCjhvWDhscEpjMTNDR3JlQ2FPM3VUNG81OVVnZmp2QVovQ0R1Yi9IMTVHL0tlMW1XZXltaTlZTWIxRU9BbXVXdHkKOHBTSlowd2p0QjUrQVM5Q0FTMTlsWGZDYm9jemVCZDJydnFURzl1YVhmVjRwQkJXRVA1NEwxMUdScWpJc2o3TwpmT2l1Tm1DUjdGZXlWV0FSVU5PWFgzOE5VSjBlWUswbjVwam1rdk5jWVVJVVdlbTZDSzR3bSsrWkFnTUJBQUdqCmdZWXdnWU13RGdZRFZSMFBBUUgvQkFRREFnV2dNQjBHQTFVZEpRUVdNQlFHQ0NzR0FRVUZCd01CQmdnckJnRUYKQlFjREFqQU1CZ05WSFJNQkFmOEVBakFBTUI4R0ExVWRJd1FZTUJhQUZMbms0MTBHRnRDL3pNS2Q3bEtrZEJibgorcE9tTUNNR0ExVWRFUVFjTUJxQ0dDb3VhSFZpWW14bExYSmxiR0Y1TG1OcGJHbDFiUzVwYnpBTkJna3Foa2lHCjl3MEJBUXNGQUFPQ0FRRUFtWjBvOFRsWG1GWGJ5YUtMamVmd1hVc0NEN1pDR3JSbkV3aUdVT1hTZlJYUTJIa0cKMDM2YUd3OWZNTHVoOXhTNVZ1L1RGM1BCZG1MV0tPckJldUdCYWNEdUlNNElLUlBIRjdXNUhhcXZxNVEybWwrdQpTVzU4UG5teno0OXRCUHNjb0RoNXRLdy9Eb0xZRUFYU1dzamVaSVp4QjNPbUtySTBNV3VONnlQNU9xRXJ2TG84CmxqU2pxd2NDMVBBcFBvWUZjVklvZWFYVktla0VzMXRZemZ4a1pOS09KaWFvSEdUbnQxNnRRK0pmZmVYdUVQVTQKUlZjektTNVJkUDYzODVCT29JWE9PR3FkbnBsTTlDNXViU2pUWit5aGhHd2UrWDgyRzczNmxFZTgxL2ZnM05KZgo1VnZQeFBINWNSUnNqSFVRV0ptcU5DMjFQWEk4YzVxUUlDYyt2dz09Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K
  tls.key: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFb3dJQkFBS0NBUUVBeEF0RitRSUwzZzBhNnZyc2RTZ0tjWlU0ODU1MW5BNzV6WkRtaG90anl4eXp0eWw2Cjlxb095WGVrVnVFUU1tRGFsRXRWZGdUVFlXaTN5MkVkbUNXY1p4SnkweHFRUmJmU1FwZVF5Qm5TY3h0cDdkeVUKRm1OQWQrdWhTaGhJeW1ZSGlzRXJIenhOdFllcnBXUG9TVWdjL1ZrS3lmQTB6dktGL0phU1hOZHdocTNnbWp0NwprK0tPZlZJSDQ3d0dmd2c3bS94OWVSdnludFpsbnNwb3ZXREc5UkRnSnJscmN2S1VpV2RNSTdRZWZnRXZRZ0V0CmZaVjN3bTZITTNnWGRxNzZreHZibWwzMWVLUVFWaEQrZUM5ZFJrYW95TEkrem56b3JqWmdrZXhYc2xWZ0VWRFQKbDE5L0RWQ2RIbUN0SithWTVwTHpYR0ZDRkZucHVnaXVNSnZ2bVFJREFRQUJBb0lCQUNlSW5uTzRsZXdSZUh3cQpYY1RDYmxpNVh1TEI4YldtejNsRTN6Z0NvLzB4ckl3ak1Vak13bTZlVWVXelBURHJseWlRaUl0a0xieFhBYmxoCnZEWVNYNWZwZ0g1UnZRWlNLM1NDWEEvK0pSSlJWT3RDc0JwVHFZeUZWK0U1UkhTTVhyajhlMVd4TTNxSUFYVTQKMEg3MnErSHJNdUhHTHVBTXlEaEhwUHhUOWIweXVHSy8wbWVZV2NHNkdjY1JPWVdxSEJxWENkSEdlRS93eU0rOQpCa0tSa3lmbEdYMEh6K2E3T01oWUE2alRBMmYwQzluOEdCY2ZYZEI4VktjMDhMUXZkc2JIbjlSZEpYN1BMN0d2CmR3ZHJEZ3RYRFByemRZaFRvYS9iWHVRRVJWamgrZXJDanRhTVFvZnBJb0lIYldiL01HRk1paVM5TGM0TnlSZW0KVnNPUFFBRUNnWUVBK0ZxWnU5SVpVSlVsa0tSZzI0MS9vVE81dndxQXgzZmpWS1laWVVuN3crcWxWbDR3R0F4dgpzbkxEK0VEeThHQ1BtZmVVUHhacWlLNnZpQjhMdVVWbm1CY0R0aHFEWnlGd3dJV0xvOEtTTmN1ZXRWSmRZUTZjCnB2bHptQm1SYmtBWWFsclZIQWxiU1plVTFhWDFJME9yZm9BckhZUlppWHVwT1RxbkpKS0tZOEVDZ1lFQXloUmsKbC9GT2ZBSFNySWhKTGpnZ2J1MG9uK054Y1dHbzB2blJkSTJXTGUvcWxPWkRwc3lTWTVHcDZkL1dCSzZudnl0SApZODJiOWpWVTV0WkFnOU90Y3NvNkhFZUJhbmVtOFpScFNzVVI1Qm5jS3VhRVJwWXJFS3VGSWIyM3RXNy9PdVBOCkNyK3pKRlZHYk1GZTE3YTgwZDRZNUdLOFhBVkFEQ0Y3NllEUm9ka0NnWUFsZkdkZjlpSmtDMThVS2Z1RXFDTHYKamdNblZzcUJVUk03SDZjTkRFRzRISjdBSG85YjBlUzZKcUIxeERmbkdHd1ViVTR2QjQ4aytsajhUdE5TTDZ1bgpSVElHTnBKRzZzRStEZW81MlpDQUZpL3Fabmc1d2g3YkJTUHhmVXA2UGFweHd5d1BnMG9JSFowVmNtdEIyMkR6ClF4MCs0MDh3ZFQzaHFYeTVCSFZuQVFLQmdFT1MzQ3grOGFyQUJVM1NhUDQrb0lIWFpqMUpGaGMrKy9CSXY0VEYKRDlJZXB3ZlJsQS9EMnJQVzhzV1ZKd0Q3MG5ZM3A3QzFBWkVzTms3V21FNDh5NFJXSVdaeGR0SStYcUhyNmVXcAp6cGpER1A5emhBb0NqellNMVFENmF1TU4wZVZFWmIxUmF6c2NGT2VySmVibVlXK2dZQnlHODh1bHFjd2txa1hqCjRMWEpBb0dCQU1VaURuYVJhd2V3UVgzOGErbGtPU3ZOOHFHSUlCVXlJMUF3amcrakJ2KzFCem1LWGtDdjVudngKbE4wdVhUUlFlY3NPZkxxZDRYSnFaMWppcC9Uem80MnJhN2R6aGNVRmo4cWJnTkJqUEo1TXJUZFpaT0s1WmZNaApVT0RzSUxySEUwdThRVnVEeHBzc01pQnJiRlEvc01XMVhuYzVKSzltS24zRFczVnk4amZZCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==
---
# Source: cilium/templates/hubble/tls-helm/server-secret.yaml
apiVersion: v1
kind: Secret
metadata:
  name: hubble-server-certs
  namespace: kube-system
type: kubernetes.io/tls
data:
  ca.crt:  LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURGRENDQWZ5Z0F3SUJBZ0lSQUxIeHRLQ2RWaUR3cU1JS1JMY1NCR0V3RFFZSktvWklodmNOQVFFTEJRQXcKRkRFU01CQUdBMVVFQXhNSlEybHNhWFZ0SUVOQk1CNFhEVEl6TURNeE56QTVNVE0xTTFvWERUSTJNRE14TmpBNQpNVE0xTTFvd0ZERVNNQkFHQTFVRUF4TUpRMmxzYVhWdElFTkJNSUlCSWpBTkJna3Foa2lHOXcwQkFRRUZBQU9DCkFROEFNSUlCQ2dLQ0FRRUF3VFJFeXJldnlsMm9XSFZaYnkvMXBmUXFscXVTVTlyV0dUMUxOd3NNcE96UlFIZi8KejlLRHZrTnVGSXZDZ2tiS0RPY1J2UkVvTFlsZGpjWGZwTFRrbFlUYnBjRFJReHd3elRXK3hlWnZSV2tndVpQUQo3WENZV3RrZ2FXcUQ5bjduMFRSU3FMdTU5TEo4L3BoYlYvYlVHSGZ0bGxEUVZCbXZ5N0Y2S3QzdDJRamw2aFNNCmY0VUhwZjFpMEdWb21qUzZWeDZvMmQxeFhNUEh6ZGVoYjQ4Wk1QYzA5UGQySUhCYlRsMk1wZjF6Z05Pb3l6MnYKZUwrSlIzV0FFY3lySVNBRU02bGtnRzZ6blEzSnV2Z0xwcHBtZ1pMbEYzVXlDYXgxTVNGMGkxb1RLNEtIWDQyaApsT3RkUyt2REl1VWVUeFRjUWxyRmdMT0hBR3RLMDQ2aTQrMVZuUUlEQVFBQm8yRXdYekFPQmdOVkhROEJBZjhFCkJBTUNBcVF3SFFZRFZSMGxCQll3RkFZSUt3WUJCUVVIQXdFR0NDc0dBUVVGQndNQ01BOEdBMVVkRXdFQi93UUYKTUFNQkFmOHdIUVlEVlIwT0JCWUVGTG5rNDEwR0Z0Qy96TUtkN2xLa2RCYm4rcE9tTUEwR0NTcUdTSWIzRFFFQgpDd1VBQTRJQkFRQW5KZHhRZXliRFpxeGQxYllMazZaeDRkYVo2Tkc3eDlyZ2lUMWgyT21tOFBFUE5OZngwa090CnZ0L0xpSGltUTlxMDl2c0pleERLQU1vL1h6Z3JxZ0pvbkNrY1RVTk1vdlVrU2loM2IxSVdSYXFNNUtRemxsN1QKbnRwem5SQXZyTElGamNUYlpIY2JMYk5qVkprY2xwWFdTYTFkVnlMOElKbVFoaUljSnQyS041TFNPdnRRbFdKTwprWFNtdVEwbUFoTW0wVEVyWnM3dWRRQ2I5RDFrd0svQlQwVU84Q2dheEZwZEZET2g0TGJ5NERMN0dnc25RY1ZGCld4UHl6elExdHdDRkxPa2dxYnVhZ25TWUdYVW9XbjRNMzREVE1USUk0N1ZNVDgzYzR6Q2hoU1RTanNLRDFxaCsKQWthUTlOazE1L01QUjExRkhVKytTeld4eGdEK0hkanoKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=
  tls.crt: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURWakNDQWo2Z0F3SUJBZ0lRSTJhbjBLeHlQTzVPRy92MGdMZ3RQakFOQmdrcWhraUc5dzBCQVFzRkFEQVUKTVJJd0VBWURWUVFERXdsRGFXeHBkVzBnUTBFd0hoY05Nak13TXpFM01Ea3hNelV6V2hjTk1qWXdNekUyTURreApNelV6V2pBcU1TZ3dKZ1lEVlFRRERCOHFMbVJsWm1GMWJIUXVhSFZpWW14bExXZHljR011WTJsc2FYVnRMbWx2Ck1JSUJJakFOQmdrcWhraUc5dzBCQVFFRkFBT0NBUThBTUlJQkNnS0NBUUVBNDVtNGk3dmtHUVdOR0RnVDEwYWMKbGprbHVEVk9NN2NvbkJ5UnBVY0JZVnpidThMY2dPVWluT3drc2FBZTZXUm9udkthSTRjc1FndVZWUnpaaVpoLwpTdE5XZWlhZURPdjNxZXRibHU1cEVtZ2FZeldwSXNJU0VJa3VRNnh6UnlGbGo4Q0pVZ1hFczQwQ0gzeVdmN2x1CkFWaXJVdTNqMkxTVUY3bjNwTDRvVHllT1hjY2xiZ1lodS8zUS9NN2pCMkwvOElrVkFqWTJoOVB0Y2RUMXpvRUoKWUFaRG9SeG1vNzVHelhVWE1MOUNjelpFRUp3eStKWDFTb1Ftc01nYU9GWVF0RzRhUVg3clUyWlNMMG81dFlucwp0ZENUa3k1ZHpLZW0zK3ZEcmVkWUUwNXNNMngyajUxb0w1Qm51YmZtaTIvVWh0UHlSSyt1SGc2cGdqcTA4dXhVCkR3SURBUUFCbzRHTk1JR0tNQTRHQTFVZER3RUIvd1FFQXdJRm9EQWRCZ05WSFNVRUZqQVVCZ2dyQmdFRkJRY0QKQVFZSUt3WUJCUVVIQXdJd0RBWURWUjBUQVFIL0JBSXdBREFmQmdOVkhTTUVHREFXZ0JTNTVPTmRCaGJRdjh6QwpuZTVTcEhRVzUvcVRwakFxQmdOVkhSRUVJekFoZ2g4cUxtUmxabUYxYkhRdWFIVmlZbXhsTFdkeWNHTXVZMmxzCmFYVnRMbWx2TUEwR0NTcUdTSWIzRFFFQkN3VUFBNElCQVFDMW50QSs5a1BPRS9XUjEyMEw2OFF4aXVJTkFMbnIKWVB6ckU4V211Mno3TGM1a0JFUzFSUG1WQ2pBRFhRb1FIK3Rrc3FzVmZTQk5oRHptak1yRUdvUEU3VTVVTkhURwphREJKQVZNRXkvZnBicHczYjRXaGduMEVwYko5N25kYmpqWTRmazFza0ZiZjZBdDViRHNSQU9Kb3dhQ05Yei9oCmVSZ2NVMmFwZWN2UkUrMXJZQmZDSG5LV1loaDZQazUrM2ZlMFNHNlNveTB5RWNUQjVKZk4zdDhJYVhPejBCazYKTHdaWlJwTERpdlBMTE80QUYrbGl3cjYyM1JwQmV3eVNKVDl3YXVLQXlxTnd1YmwyRkxNeEs3ZFI1RUllS3ZZbQp6SDIvdDhoaHNrRWpOVjFmTkcxYVZxWnpjaWQ4K09KYW96R3JzWWM3YkJ1Z0R3VTRyWmg5L1RkTwotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==
  tls.key: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcFFJQkFBS0NBUUVBNDVtNGk3dmtHUVdOR0RnVDEwYWNsamtsdURWT003Y29uQnlScFVjQllWemJ1OExjCmdPVWluT3drc2FBZTZXUm9udkthSTRjc1FndVZWUnpaaVpoL1N0TldlaWFlRE92M3FldGJsdTVwRW1nYVl6V3AKSXNJU0VJa3VRNnh6UnlGbGo4Q0pVZ1hFczQwQ0gzeVdmN2x1QVZpclV1M2oyTFNVRjduM3BMNG9UeWVPWGNjbApiZ1lodS8zUS9NN2pCMkwvOElrVkFqWTJoOVB0Y2RUMXpvRUpZQVpEb1J4bW83NUd6WFVYTUw5Q2N6WkVFSnd5CitKWDFTb1Ftc01nYU9GWVF0RzRhUVg3clUyWlNMMG81dFluc3RkQ1RreTVkektlbTMrdkRyZWRZRTA1c00yeDIKajUxb0w1Qm51YmZtaTIvVWh0UHlSSyt1SGc2cGdqcTA4dXhVRHdJREFRQUJBb0lCQVFEVlFmRHdSVEpaMjZDego0Nzk3Zy9jdXJ3ZFB6ZXFqZkdmNXRxOGxmUjJtKzlvTDJXN0ErM0h1TlRuQWtYYkZXMGJJUUNyN1JTbk1ESXgwCi9wNDZWZ0JYdlNRWE9sMzNYNVpreVZtOVYxQnVaY3JyMEpqVkw2QzdpNzRrdk00YkJRampwQlZISEk2TmFuOWwKdjFoSS94YzYvYmt4OENNQXAxcm56R0ZsSktRaVhQcW04Mi9xdmFLUHltZmx3alhQdWZKeXBDRThqY3dPNFhKRwpGbFdxQk5tMWdqWVhnZVQ3VytuY1JvYzNkNHRjbnZtWlJMRnB1ZUZ1RWx3dllWY3plZUUySmVrUTlOSlhnRXJMCjdTaE90NThhekhlb1ZVcTVPcEY4QUU5VERJNVhDZ0NoclhudDYzVkEwTnpWbkkrbkszUlI5ZnZXdmxNbzlhWGQKNGhoNjVCdkJBb0dCQVBCYWNoWm92Yk5YVFFMODk4elV2MytjR2FJOFJuN2E0Qm9tK003V3RySkt3b2wyQXFlZQpUN3BvU1JXWWowbllLY0J2VTAzSFhsRlNzRTE5UXN2Zy8wKzZBQzlCL0xaVmtqc1JrbDZSbW42SmtKUm0ybUtUClBwZUtFVk94V25oZWF4QmNWOTJIaFdsUTRkZ3RicC8wS0x3Uk5VVWZNYVlieUFIUFJWL0c3NnJoQW9HQkFQSnEKdnVxRVZ5MzNISEJVNnpRVzBRQUFCMzQxdG9SOXVaZnJ1aWZrVngzdUZ2V0JiUmprdUFkVXpxekZELzhPOWJGVApDcUFyaStHRW03SkY2V2hJOStSRWxkbmNPa2FFS1JXdmEvQnp6SjdvNXpGb3lCUzZPbllJc3orVmd4bXVXN0lTCk5qYWJORmZxb2piQnRIY3VQdkJxVm1hZlNVVWZ0WkxlZGdWcGtVenZBb0dBVGNTbUIzUXFkUjI1TUU5VGluWUgKNUMxSTZnSmd1T2p1KytkQ09BS25LSGNpRE1JZlI4YmtleWNGQnJUUElCQ09LZEtiZko0V2VXK3MxZFhDeUI3cgozUXNNeGoydW0veUNEUlM1YkZubVNDMFFsOFBUdzNOckhETXpPZ1kzaEp6Z1BYSHppQjB5WUlvb0dQOVNQUFVPClBSUEFUYll6SlZEMTNRZ0lwVjNEN0dFQ2dZRUEzN25zb1B1cWlkMTUvYUlod0YwZVhtV29oSzZGMkJsQVpCbEcKSVBMNEE4TnNwUC9oOUF1Q1hDSEU3R2Fpc0w3WnVlSHQrSXkzK0ZZdWE0VmlPTUMvSjRpMDAvQVFTR3hJanA3cgplMnNqK2JUeFNnUnVROUxyaVd2V0ltU1dMZWxnN3lNbnJaWG41UXZDMGM1TUE0Skd6Qk1YMG5aSFpPZ3k1MjB3CmR5WksxemtDZ1lFQXVmWkhXY0M1eHkzczJiR2VWS3RMOWZaZFptb2JNL2FINHIzVmdnL2JsZlV3S3VVa0hpanoKa1FRWWZlcGZMWW96Y3hQRWFtSTlXT2NqamZEeDNLSCtTOEdVb2V1Wmtsb0dTTWZ5NW5lR0Q0M2NlWWlHMTNnZQpyOFBKNjArUjE0b0xUVUFUYWhwUkdQaEVNZ2o4M1cvMTR5NElPNDJFU25SOXhaUkI3MDUrWm8wPQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo=
---
# Source: cilium/templates/cilium-configmap.yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: cilium-config
  namespace: kube-system
data:

  # Identity allocation mode selects how identities are shared between cilium
  # nodes by setting how they are stored. The options are "crd" or "kvstore".
  # - "crd" stores identities in kubernetes as CRDs (custom resource definition).
  #   These can be queried with:
  #     kubectl get ciliumid
  # - "kvstore" stores identities in an etcd kvstore, that is
  #   configured below. Cilium versions before 1.6 supported only the kvstore
  #   backend. Upgrades from these older cilium versions should continue using
  #   the kvstore by commenting out the identity-allocation-mode below, or
  #   setting it to "kvstore".
  identity-allocation-mode: crd
  identity-heartbeat-timeout: "30m0s"
  identity-gc-interval: "15m0s"
  cilium-endpoint-gc-interval: "5m0s"
  nodes-gc-interval: "5m0s"
  skip-cnp-status-startup-clean: "false"
  # Disable the usage of CiliumEndpoint CRD
  disable-endpoint-crd: "false"

  # If you want to run cilium in debug mode change this value to true
  debug: "false"
  debug-verbose: ""
  # The agent can be put into the following three policy enforcement modes
  # default, always and never.
  # https://docs.cilium.io/en/latest/security/policy/intro/#policy-enforcement-modes
  enable-policy: "default"

  # Enable IPv4 addressing. If enabled, all endpoints are allocated an IPv4
  # address.
  enable-ipv4: "true"

  # Enable IPv6 addressing. If enabled, all endpoints are allocated an IPv6
  # address.
  enable-ipv6: "false"
  # Users who wish to specify their own custom CNI configuration file must set
  # custom-cni-conf to "true", otherwise Cilium may overwrite the configuration.
  custom-cni-conf: "false"
  enable-bpf-clock-probe: "true"
  # If you want cilium monitor to aggregate tracing for packets, set this level
  # to "low", "medium", or "maximum". The higher the level, the less packets
  # that will be seen in monitor output.
  monitor-aggregation: medium

  # The monitor aggregation interval governs the typical time between monitor
  # notification events for each allowed connection.
  #
  # Only effective when monitor aggregation is set to "medium" or higher.
  monitor-aggregation-interval: "5s"

  # The monitor aggregation flags determine which TCP flags which, upon the
  # first observation, cause monitor notifications to be generated.
  #
  # Only effective when monitor aggregation is set to "medium" or higher.
  monitor-aggregation-flags: all
  # Specifies the ratio (0.0-1.0] of total system memory to use for dynamic
  # sizing of the TCP CT, non-TCP CT, NAT and policy BPF maps.
  bpf-map-dynamic-size-ratio: "0.0025"
  # bpf-policy-map-max specifies the maximum number of entries in endpoint
  # policy map (per endpoint)
  bpf-policy-map-max: "16384"
  # bpf-lb-map-max specifies the maximum number of entries in bpf lb service,
  # backend and affinity maps.
  bpf-lb-map-max: "65536"
  bpf-lb-external-clusterip: "false"

  # Pre-allocation of map entries allows per-packet latency to be reduced, at
  # the expense of up-front memory allocation for the entries in the maps. The
  # default value below will minimize memory usage in the default installation;
  # users who are sensitive to latency may consider setting this to "true".
  #
  # This option was introduced in Cilium 1.4. Cilium 1.3 and earlier ignore
  # this option and behave as though it is set to "true".
  #
  # If this value is modified, then during the next Cilium startup the restore
  # of existing endpoints and tracking of ongoing connections may be disrupted.
  # As a result, reply packets may be dropped and the load-balancing decisions
  # for established connections may change.
  #
  # If this option is set to "false" during an upgrade from 1.3 or earlier to
  # 1.4 or later, then it may cause one-time disruptions during the upgrade.
  preallocate-bpf-maps: "false"

  # Regular expression matching compatible Istio sidecar istio-proxy
  # container image names
  sidecar-istio-proxy-image: "cilium/istio_proxy"

  # Name of the cluster. Only relevant when building a mesh of clusters.
  cluster-name: default
  # Unique ID of the cluster. Must be unique across all conneted clusters and
  # in the range of 1 and 255. Only relevant when building a mesh of clusters.
  cluster-id: "0"

  # Encapsulation mode for communication between nodes
  # Possible values:
  #   - disabled
  #   - vxlan (default)
  #   - geneve
  tunnel: "vxlan"


  # Enables L7 proxy for L7 policy enforcement and visibility
  enable-l7-proxy: "true"

  enable-ipv4-masquerade: "true"
  enable-ipv6-big-tcp: "false"
  enable-ipv6-masquerade: "true"

  enable-xt-socket-fallback: "true"
  install-iptables-rules: "true"
  install-no-conntrack-iptables-rules: "false"

  auto-direct-node-routes: "false"
  enable-local-redirect-policy: "false"

  kube-proxy-replacement: "strict"
  kube-proxy-replacement-healthz-bind-address: ""
  bpf-lb-sock: "false"
  enable-health-check-nodeport: "true"
  node-port-bind-protection: "true"
  enable-auto-protect-node-port-range: "true"
  enable-svc-source-range-check: "true"
  enable-l2-neigh-discovery: "true"
  arping-refresh-period: "30s"
  enable-endpoint-health-checking: "true"
  enable-health-checking: "true"
  enable-well-known-identities: "false"
  enable-remote-node-identity: "true"
  synchronize-k8s-nodes: "true"
  operator-api-serve-addr: "127.0.0.1:9234"
  # Enable Hubble gRPC service.
  enable-hubble: "true"
  # UNIX domain socket for Hubble server to listen to.
  hubble-socket-path: "/var/run/cilium/hubble.sock"
  # An additional address for Hubble server to listen to (e.g. ":4244").
  hubble-listen-address: ":4244"
  hubble-disable-tls: "false"
  hubble-tls-cert-file: /var/lib/cilium/tls/hubble/server.crt
  hubble-tls-key-file: /var/lib/cilium/tls/hubble/server.key
  hubble-tls-client-ca-files: /var/lib/cilium/tls/hubble/client-ca.crt
  ipam: "cluster-pool"
  cluster-pool-ipv4-cidr: "10.0.0.0/8"
  cluster-pool-ipv4-mask-size: "24"
  disable-cnp-status-updates: "true"
  enable-vtep: "false"
  vtep-endpoint: ""
  vtep-cidr: ""
  vtep-mask: ""
  vtep-mac: ""
  enable-bgp-control-plane: "false"
  procfs: "/host/proc"
  bpf-root: "/sys/fs/bpf"
  cgroup-root: "/run/cilium/cgroupv2"
  enable-k8s-terminating-endpoint: "true"
  enable-sctp: "false"
  remove-cilium-node-taints: "true"
  set-cilium-is-up-condition: "true"
  unmanaged-pod-watcher-interval: "15"
  tofqdns-dns-reject-response-code: "refused"
  tofqdns-enable-dns-compression: "true"
  tofqdns-endpoint-max-ip-per-hostname: "50"
  tofqdns-idle-connection-grace-period: "0s"
  tofqdns-max-deferred-connection-deletes: "10000"
  tofqdns-min-ttl: "3600"
  tofqdns-proxy-response-max-delay: "100ms"
  agent-not-ready-taint-key: "node.cilium.io/agent-not-ready"
---
# Source: cilium/templates/hubble-relay/configmap.yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: hubble-relay-config
  namespace: kube-system
data:
  config.yaml: |
    cluster-name: default
    peer-service: "hubble-peer.kube-system.svc.cluster.local:443"
    listen-address: :4245
    dial-timeout: 
    retry-timeout: 
    sort-buffer-len-max: 
    sort-buffer-drain-timeout: 
    tls-client-cert-file: /var/lib/hubble-relay/tls/client.crt
    tls-client-key-file: /var/lib/hubble-relay/tls/client.key
    tls-hubble-server-ca-files: /var/lib/hubble-relay/tls/hubble-server-ca.crt
    disable-server-tls: true
---
# Source: cilium/templates/hubble-ui/configmap.yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: hubble-ui-nginx
  namespace: kube-system
data:
  nginx.conf: "server {\n    listen       8081;\n    listen       [::]:8081;\n    server_name  localhost;\n    root /app;\n    index index.html;\n    client_max_body_size 1G;\n\n    location / {\n        proxy_set_header Host $host;\n        proxy_set_header X-Real-IP $remote_addr;\n\n        # CORS\n        add_header Access-Control-Allow-Methods \"GET, POST, PUT, HEAD, DELETE, OPTIONS\";\n        add_header Access-Control-Allow-Origin *;\n        add_header Access-Control-Max-Age 1728000;\n        add_header Access-Control-Expose-Headers content-length,grpc-status,grpc-message;\n        add_header Access-Control-Allow-Headers range,keep-alive,user-agent,cache-control,content-type,content-transfer-encoding,x-accept-content-transfer-encoding,x-accept-response-streaming,x-user-agent,x-grpc-web,grpc-timeout;\n        if ($request_method = OPTIONS) {\n            return 204;\n        }\n        # /CORS\n\n        location /api {\n            proxy_http_version 1.1;\n            proxy_pass_request_headers on;\n            proxy_hide_header Access-Control-Allow-Origin;\n            proxy_pass http://127.0.0.1:8090;\n        }\n\n        location / {\n            try_files $uri $uri/ /index.html;\n        }\n    }\n}"
---
# Source: cilium/templates/cilium-agent/clusterrole.yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: cilium
  labels:
    app.kubernetes.io/part-of: cilium
rules:
- apiGroups:
  - networking.k8s.io
  resources:
  - networkpolicies
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - discovery.k8s.io
  resources:
  - endpointslices
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - ""
  resources:
  - namespaces
  - services
  - pods
  - endpoints
  - nodes
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - apiextensions.k8s.io
  resources:
  - customresourcedefinitions
  verbs:
  - list
  - watch
  # This is used when validating policies in preflight. This will need to stay
  # until we figure out how to avoid "get" inside the preflight, and then
  # should be removed ideally.
  - get
- apiGroups:
  - cilium.io
  resources:
  - ciliumloadbalancerippools
  - ciliumbgppeeringpolicies
  - ciliumclusterwideenvoyconfigs
  - ciliumclusterwidenetworkpolicies
  - ciliumegressgatewaypolicies
  - ciliumendpoints
  - ciliumendpointslices
  - ciliumenvoyconfigs
  - ciliumidentities
  - ciliumlocalredirectpolicies
  - ciliumnetworkpolicies
  - ciliumnodes
  - ciliumnodeconfigs
  verbs:
  - list
  - watch
- apiGroups:
  - cilium.io
  resources:
  - ciliumidentities
  - ciliumendpoints
  - ciliumnodes
  verbs:
  - create
- apiGroups:
  - cilium.io
  # To synchronize garbage collection of such resources
  resources:
  - ciliumidentities
  verbs:
  - update
- apiGroups:
  - cilium.io
  resources:
  - ciliumendpoints
  verbs:
  - delete
  - get
- apiGroups:
  - cilium.io
  resources:
  - ciliumnodes
  - ciliumnodes/status
  verbs:
  - get
  - update
- apiGroups:
  - cilium.io
  resources:
  - ciliumnetworkpolicies/status
  - ciliumclusterwidenetworkpolicies/status
  - ciliumendpoints/status
  - ciliumendpoints
  verbs:
  - patch
---
# Source: cilium/templates/cilium-operator/clusterrole.yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: cilium-operator
  labels:
    app.kubernetes.io/part-of: cilium
rules:
- apiGroups:
  - ""
  resources:
  - pods
  verbs:
  - get
  - list
  - watch
  # to automatically delete [core|kube]dns pods so that are starting to being
  # managed by Cilium
  - delete
- apiGroups:
  - ""
  resources:
  - nodes
  verbs:
  - list
  - watch
- apiGroups:
  - ""
  resources:
  # To remove node taints
  - nodes
  # To set NetworkUnavailable false on startup
  - nodes/status
  verbs:
  - patch
- apiGroups:
  - discovery.k8s.io
  resources:
  - endpointslices
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - ""
  resources:
  # to perform LB IP allocation for BGP
  - services/status
  verbs:
  - update
  - patch
- apiGroups:
  - ""
  resources:
  # to check apiserver connectivity
  - namespaces
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - ""
  resources:
  - services
  - endpoints
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - cilium.io
  resources:
  - ciliumnetworkpolicies
  - ciliumclusterwidenetworkpolicies
  verbs:
  # Create auto-generated CNPs and CCNPs from Policies that have 'toGroups'
  - create
  - update
  - deletecollection
  # To update the status of the CNPs and CCNPs
  - patch
  - get
  - list
  - watch
- apiGroups:
  - cilium.io
  resources:
  - ciliumnetworkpolicies/status
  - ciliumclusterwidenetworkpolicies/status
  verbs:
  # Update the auto-generated CNPs and CCNPs status.
  - patch
  - update
- apiGroups:
  - cilium.io
  resources:
  - ciliumendpoints
  - ciliumidentities
  verbs:
  # To perform garbage collection of such resources
  - delete
  - list
  - watch
- apiGroups:
  - cilium.io
  resources:
  - ciliumidentities
  verbs:
  # To synchronize garbage collection of such resources
  - update
- apiGroups:
  - cilium.io
  resources:
  - ciliumnodes
  verbs:
  - create
  - update
  - get
  - list
  - watch
    # To perform CiliumNode garbage collector
  - delete
- apiGroups:
  - cilium.io
  resources:
  - ciliumnodes/status
  verbs:
  - update
- apiGroups:
  - cilium.io
  resources:
  - ciliumendpointslices
  - ciliumenvoyconfigs
  verbs:
  - create
  - update
  - get
  - list
  - watch
  - delete
  - patch
- apiGroups:
  - apiextensions.k8s.io
  resources:
  - customresourcedefinitions
  verbs:
  - create
  - get
  - list
  - watch
- apiGroups:
  - apiextensions.k8s.io
  resources:
  - customresourcedefinitions
  verbs:
  - update
  resourceNames:
  - ciliumloadbalancerippools.cilium.io
  - ciliumbgppeeringpolicies.cilium.io
  - ciliumclusterwideenvoyconfigs.cilium.io
  - ciliumclusterwidenetworkpolicies.cilium.io
  - ciliumegressgatewaypolicies.cilium.io
  - ciliumendpoints.cilium.io
  - ciliumendpointslices.cilium.io
  - ciliumenvoyconfigs.cilium.io
  - ciliumexternalworkloads.cilium.io
  - ciliumidentities.cilium.io
  - ciliumlocalredirectpolicies.cilium.io
  - ciliumnetworkpolicies.cilium.io
  - ciliumnodes.cilium.io
  - ciliumnodeconfigs.cilium.io
- apiGroups:
  - cilium.io
  resources:
  - ciliumloadbalancerippools
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - cilium.io
  resources:
  - ciliumloadbalancerippools/status
  verbs:
  - patch
# For cilium-operator running in HA mode.
#
# Cilium operator running in HA mode requires the use of ResourceLock for Leader Election
# between multiple running instances.
# The preferred way of doing this is to use LeasesResourceLock as edits to Leases are less
# common and fewer objects in the cluster watch "all Leases".
- apiGroups:
  - coordination.k8s.io
  resources:
  - leases
  verbs:
  - create
  - get
  - update
---
# Source: cilium/templates/hubble-ui/clusterrole.yaml
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: hubble-ui
  labels:
    app.kubernetes.io/part-of: cilium
rules:
- apiGroups:
  - networking.k8s.io
  resources:
  - networkpolicies
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - ""
  resources:
  - componentstatuses
  - endpoints
  - namespaces
  - nodes
  - pods
  - services
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - apiextensions.k8s.io
  resources:
  - customresourcedefinitions
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - cilium.io
  resources:
  - "*"
  verbs:
  - get
  - list
  - watch
---
# Source: cilium/templates/cilium-agent/clusterrolebinding.yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: cilium
  labels:
    app.kubernetes.io/part-of: cilium
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cilium
subjects:
- kind: ServiceAccount
  name: "cilium"
  namespace: kube-system
---
# Source: cilium/templates/cilium-operator/clusterrolebinding.yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: cilium-operator
  labels:
    app.kubernetes.io/part-of: cilium
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cilium-operator
subjects:
- kind: ServiceAccount
  name: "cilium-operator"
  namespace: kube-system
---
# Source: cilium/templates/hubble-ui/clusterrolebinding.yaml
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: hubble-ui
  labels:
    app.kubernetes.io/part-of: cilium
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: hubble-ui
subjects:
- kind: ServiceAccount
  name: "hubble-ui"
  namespace: kube-system
---
# Source: cilium/templates/cilium-agent/role.yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: cilium-config-agent
  namespace: kube-system
  labels:
    app.kubernetes.io/part-of: cilium
rules:
- apiGroups:
  - ""
  resources:
  - configmaps
  verbs:
  - get
  - list
  - watch
---
# Source: cilium/templates/cilium-agent/rolebinding.yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: cilium-config-agent
  namespace: kube-system
  labels:
    app.kubernetes.io/part-of: cilium
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: cilium-config-agent
subjects:
  - kind: ServiceAccount
    name: "cilium"
    namespace: kube-system
---
# Source: cilium/templates/hubble-relay/service.yaml
kind: Service
apiVersion: v1
metadata:
  name: hubble-relay
  namespace: kube-system
  labels:
    k8s-app: hubble-relay
    app.kubernetes.io/name: hubble-relay
    app.kubernetes.io/part-of: cilium
spec:
  type: "ClusterIP"
  selector:
    k8s-app: hubble-relay
  ports:
  - protocol: TCP
    port: 80
    targetPort: 4245
---
# Source: cilium/templates/hubble-ui/service.yaml
kind: Service
apiVersion: v1
metadata:
  name: hubble-ui
  namespace: kube-system
  labels:
    k8s-app: hubble-ui
    app.kubernetes.io/name: hubble-ui
    app.kubernetes.io/part-of: cilium
spec:
  type: "ClusterIP"
  selector:
    k8s-app: hubble-ui
  ports:
    - name: http
      port: 80
      targetPort: 8081
---
# Source: cilium/templates/hubble/peer-service.yaml
apiVersion: v1
kind: Service
metadata:
  name: hubble-peer
  namespace: kube-system
  labels:
    k8s-app: cilium
    app.kubernetes.io/part-of: cilium
    app.kubernetes.io/name: hubble-peer
spec:
  selector:
    k8s-app: cilium
  ports:
  - name: peer-service
    port: 443
    protocol: TCP
    targetPort: 4244
  internalTrafficPolicy: Local
---
# Source: cilium/templates/cilium-agent/daemonset.yaml
apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: cilium
  namespace: kube-system
  labels:
    k8s-app: cilium
    app.kubernetes.io/part-of: cilium
    app.kubernetes.io/name: cilium-agent
spec:
  selector:
    matchLabels:
      k8s-app: cilium
  updateStrategy:
    rollingUpdate:
      maxUnavailable: 2
    type: RollingUpdate
  template:
    metadata:
      annotations:
        # Set app AppArmor's profile to "unconfined". The value of this annotation
        # can be modified as long users know which profiles they have available
        # in AppArmor.
        container.apparmor.security.beta.kubernetes.io/cilium-agent: "unconfined"
        container.apparmor.security.beta.kubernetes.io/clean-cilium-state: "unconfined"
        container.apparmor.security.beta.kubernetes.io/mount-cgroup: "unconfined"
        container.apparmor.security.beta.kubernetes.io/apply-sysctl-overwrites: "unconfined"
      labels:
        k8s-app: cilium
        app.kubernetes.io/name: cilium-agent
        app.kubernetes.io/part-of: cilium
    spec:
      containers:
      - name: cilium-agent
        image: "quay.io/cilium/cilium:v1.13.0@sha256:6544a3441b086a2e09005d3e21d1a4afb216fae19c5a60b35793c8a9438f8f68"
        imagePullPolicy: IfNotPresent
        command:
        - cilium-agent
        args:
        - --config-dir=/tmp/cilium/config-map
        startupProbe:
          httpGet:
            host: "127.0.0.1"
            path: /healthz
            port: 9879
            scheme: HTTP
            httpHeaders:
            - name: "brief"
              value: "true"
          failureThreshold: 105
          periodSeconds: 2
          successThreshold: 1
        livenessProbe:
          httpGet:
            host: "127.0.0.1"
            path: /healthz
            port: 9879
            scheme: HTTP
            httpHeaders:
            - name: "brief"
              value: "true"
          periodSeconds: 30
          successThreshold: 1
          failureThreshold: 10
          timeoutSeconds: 5
        readinessProbe:
          httpGet:
            host: "127.0.0.1"
            path: /healthz
            port: 9879
            scheme: HTTP
            httpHeaders:
            - name: "brief"
              value: "true"
          periodSeconds: 30
          successThreshold: 1
          failureThreshold: 3
          timeoutSeconds: 5
        env:
        - name: K8S_NODE_NAME
          valueFrom:
            fieldRef:
              apiVersion: v1
              fieldPath: spec.nodeName
        - name: CILIUM_K8S_NAMESPACE
          valueFrom:
            fieldRef:
              apiVersion: v1
              fieldPath: metadata.namespace
        - name: CILIUM_CLUSTERMESH_CONFIG
          value: /var/lib/cilium/clustermesh/
        - name: CILIUM_CNI_CHAINING_MODE
          valueFrom:
            configMapKeyRef:
              name: cilium-config
              key: cni-chaining-mode
              optional: true
        - name: CILIUM_CUSTOM_CNI_CONF
          valueFrom:
            configMapKeyRef:
              name: cilium-config
              key: custom-cni-conf
              optional: true
        lifecycle:
          postStart:
            exec:
              command:
              - "/cni-install.sh"
              - "--enable-debug=false"
              - "--cni-exclusive=true"
              - "--log-file=/var/run/cilium/cilium-cni.log"
          preStop:
            exec:
              command:
              - /cni-uninstall.sh
        securityContext:
          seLinuxOptions:
            level: s0
            type: spc_t
          capabilities:
            add:
              - CHOWN
              - KILL
              - NET_ADMIN
              - NET_RAW
              - IPC_LOCK
              - SYS_MODULE
              - SYS_ADMIN
              - SYS_RESOURCE
              - DAC_OVERRIDE
              - FOWNER
              - SETGID
              - SETUID
            drop:
              - ALL
        terminationMessagePolicy: FallbackToLogsOnError
        volumeMounts:
        # Unprivileged containers need to mount /proc/sys/net from the host
        # to have write access
        - mountPath: /host/proc/sys/net
          name: host-proc-sys-net
        # Unprivileged containers need to mount /proc/sys/kernel from the host
        # to have write access
        - mountPath: /host/proc/sys/kernel
          name: host-proc-sys-kernel
        - name: bpf-maps
          mountPath: /sys/fs/bpf
          # Unprivileged containers can't set mount propagation to bidirectional
          # in this case we will mount the bpf fs from an init container that
          # is privileged and set the mount propagation from host to container
          # in Cilium.
          mountPropagation: HostToContainer
        - name: cilium-run
          mountPath: /var/run/cilium
        - name: cni-path
          mountPath: /host/opt/cni/bin
        - name: etc-cni-netd
          mountPath: /host/etc/cni/net.d
        - name: clustermesh-secrets
          mountPath: /var/lib/cilium/clustermesh
          readOnly: true
          # Needed to be able to load kernel modules
        - name: lib-modules
          mountPath: /lib/modules
          readOnly: true
        - name: xtables-lock
          mountPath: /run/xtables.lock
        - name: hubble-tls
          mountPath: /var/lib/cilium/tls/hubble
          readOnly: true
        - name: tmp
          mountPath: /tmp
      initContainers:
      - name: config
        image: "quay.io/cilium/cilium:v1.13.0@sha256:6544a3441b086a2e09005d3e21d1a4afb216fae19c5a60b35793c8a9438f8f68"
        imagePullPolicy: IfNotPresent
        command:
        - cilium
        - build-config
        env:
        - name: K8S_NODE_NAME
          valueFrom:
            fieldRef:
              apiVersion: v1
              fieldPath: spec.nodeName
        - name: CILIUM_K8S_NAMESPACE
          valueFrom:
            fieldRef:
              apiVersion: v1
              fieldPath: metadata.namespace
        volumeMounts:
        - name: tmp
          mountPath: /tmp
        terminationMessagePolicy: FallbackToLogsOnError
      # Required to mount cgroup2 filesystem on the underlying Kubernetes node.
      # We use nsenter command with host's cgroup and mount namespaces enabled.
      - name: mount-cgroup
        image: "quay.io/cilium/cilium:v1.13.0@sha256:6544a3441b086a2e09005d3e21d1a4afb216fae19c5a60b35793c8a9438f8f68"
        imagePullPolicy: IfNotPresent
        env:
        - name: CGROUP_ROOT
          value: /run/cilium/cgroupv2
        - name: BIN_PATH
          value: /opt/cni/bin
        command:
        - sh
        - -ec
        # The statically linked Go program binary is invoked to avoid any
        # dependency on utilities like sh and mount that can be missing on certain
        # distros installed on the underlying host. Copy the binary to the
        # same directory where we install cilium cni plugin so that exec permissions
        # are available.
        - |
          cp /usr/bin/cilium-mount /hostbin/cilium-mount;
          nsenter --cgroup=/hostproc/1/ns/cgroup --mount=/hostproc/1/ns/mnt "${BIN_PATH}/cilium-mount" $CGROUP_ROOT;
          rm /hostbin/cilium-mount
        volumeMounts:
        - name: hostproc
          mountPath: /hostproc
        - name: cni-path
          mountPath: /hostbin
        terminationMessagePolicy: FallbackToLogsOnError
        securityContext:
          seLinuxOptions:
            level: s0
            type: spc_t
          capabilities:
            add:
              - SYS_ADMIN
              - SYS_CHROOT
              - SYS_PTRACE
            drop:
              - ALL
      - name: apply-sysctl-overwrites
        image: "quay.io/cilium/cilium:v1.13.0@sha256:6544a3441b086a2e09005d3e21d1a4afb216fae19c5a60b35793c8a9438f8f68"
        imagePullPolicy: IfNotPresent
        env:
        - name: BIN_PATH
          value: /opt/cni/bin
        command:
        - sh
        - -ec
        # The statically linked Go program binary is invoked to avoid any
        # dependency on utilities like sh that can be missing on certain
        # distros installed on the underlying host. Copy the binary to the
        # same directory where we install cilium cni plugin so that exec permissions
        # are available.
        - |
          cp /usr/bin/cilium-sysctlfix /hostbin/cilium-sysctlfix;
          nsenter --mount=/hostproc/1/ns/mnt "${BIN_PATH}/cilium-sysctlfix";
          rm /hostbin/cilium-sysctlfix
        volumeMounts:
        - name: hostproc
          mountPath: /hostproc
        - name: cni-path
          mountPath: /hostbin
        terminationMessagePolicy: FallbackToLogsOnError
        securityContext:
          seLinuxOptions:
            level: s0
            type: spc_t
          capabilities:
            add:
              - SYS_ADMIN
              - SYS_CHROOT
              - SYS_PTRACE
            drop:
              - ALL
      # Mount the bpf fs if it is not mounted. We will perform this task
      # from a privileged container because the mount propagation bidirectional
      # only works from privileged containers.
      - name: mount-bpf-fs
        image: "quay.io/cilium/cilium:v1.13.0@sha256:6544a3441b086a2e09005d3e21d1a4afb216fae19c5a60b35793c8a9438f8f68"
        imagePullPolicy: IfNotPresent
        args:
        - 'mount | grep "/sys/fs/bpf type bpf" || mount -t bpf bpf /sys/fs/bpf'
        command:
        - /bin/bash
        - -c
        - --
        terminationMessagePolicy: FallbackToLogsOnError
        securityContext:
          privileged: true
        volumeMounts:
        - name: bpf-maps
          mountPath: /sys/fs/bpf
          mountPropagation: Bidirectional
      - name: clean-cilium-state
        image: "quay.io/cilium/cilium:v1.13.0@sha256:6544a3441b086a2e09005d3e21d1a4afb216fae19c5a60b35793c8a9438f8f68"
        imagePullPolicy: IfNotPresent
        command:
        - /init-container.sh
        env:
        - name: CILIUM_ALL_STATE
          valueFrom:
            configMapKeyRef:
              name: cilium-config
              key: clean-cilium-state
              optional: true
        - name: CILIUM_BPF_STATE
          valueFrom:
            configMapKeyRef:
              name: cilium-config
              key: clean-cilium-bpf-state
              optional: true
        terminationMessagePolicy: FallbackToLogsOnError
        securityContext:
          seLinuxOptions:
            level: s0
            type: spc_t
          capabilities:
            add:
              - NET_ADMIN
              - SYS_MODULE
              - SYS_ADMIN
              - SYS_RESOURCE
            drop:
              - ALL
        volumeMounts:
        - name: bpf-maps
          mountPath: /sys/fs/bpf
          # Required to mount cgroup filesystem from the host to cilium agent pod
        - name: cilium-cgroup
          mountPath: /run/cilium/cgroupv2
          mountPropagation: HostToContainer
        - name: cilium-run
          mountPath: /var/run/cilium
        resources:
          requests:
            cpu: 100m
            memory: 100Mi # wait-for-kube-proxy
      restartPolicy: Always
      priorityClassName: system-node-critical
      serviceAccount: "cilium"
      serviceAccountName: "cilium"
      terminationGracePeriodSeconds: 1
      hostNetwork: true
      affinity:
        podAntiAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
          - labelSelector:
              matchLabels:
                k8s-app: cilium
            topologyKey: kubernetes.io/hostname
      nodeSelector:
        kubernetes.io/os: linux
      tolerations:
        - operator: Exists
      volumes:
        # For sharing configuration between the "config" initContainer and the agent
      - name: tmp
        emptyDir: {}
        # To keep state between restarts / upgrades
      - name: cilium-run
        hostPath:
          path: /var/run/cilium
          type: DirectoryOrCreate
        # To keep state between restarts / upgrades for bpf maps
      - name: bpf-maps
        hostPath:
          path: /sys/fs/bpf
          type: DirectoryOrCreate
      # To mount cgroup2 filesystem on the host
      - name: hostproc
        hostPath:
          path: /proc
          type: Directory
      # To keep state between restarts / upgrades for cgroup2 filesystem
      - name: cilium-cgroup
        hostPath:
          path: /run/cilium/cgroupv2
          type: DirectoryOrCreate
      # To install cilium cni plugin in the host
      - name: cni-path
        hostPath:
          path:  /opt/cni/bin
          type: DirectoryOrCreate
        # To install cilium cni configuration in the host
      - name: etc-cni-netd
        hostPath:
          path: /etc/cni/net.d
          type: DirectoryOrCreate
        # To be able to load kernel modules
      - name: lib-modules
        hostPath:
          path: /lib/modules
        # To access iptables concurrently with other processes (e.g. kube-proxy)
      - name: xtables-lock
        hostPath:
          path: /run/xtables.lock
          type: FileOrCreate
        # To read the clustermesh configuration
      - name: clustermesh-secrets
        secret:
          secretName: cilium-clustermesh
          # note: the leading zero means this number is in octal representation: do not remove it
          defaultMode: 0400
          optional: true
      - name: host-proc-sys-net
        hostPath:
          path: /proc/sys/net
          type: Directory
      - name: host-proc-sys-kernel
        hostPath:
          path: /proc/sys/kernel
          type: Directory
      - name: hubble-tls
        projected:
          # note: the leading zero means this number is in octal representation: do not remove it
          defaultMode: 0400
          sources:
          - secret:
              name: hubble-server-certs
              optional: true
              items:
              - key: ca.crt
                path: client-ca.crt
              - key: tls.crt
                path: server.crt
              - key: tls.key
                path: server.key
---
# Source: cilium/templates/cilium-operator/deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: cilium-operator
  namespace: kube-system
  labels:
    io.cilium/app: operator
    name: cilium-operator
    app.kubernetes.io/part-of: cilium
    app.kubernetes.io/name: cilium-operator
spec:
  # See docs on ServerCapabilities.LeasesResourceLock in file pkg/k8s/version/version.go
  # for more details.
  replicas: 2
  selector:
    matchLabels:
      io.cilium/app: operator
      name: cilium-operator
  strategy:
    rollingUpdate:
      maxSurge: 1
      maxUnavailable: 1
    type: RollingUpdate
  template:
    metadata:
      annotations:
      labels:
        io.cilium/app: operator
        name: cilium-operator
        app.kubernetes.io/part-of: cilium
        app.kubernetes.io/name: cilium-operator
    spec:
      containers:
      - name: cilium-operator
        image: "quay.io/cilium/operator-generic:v1.13.0@sha256:4b58d5b33e53378355f6e8ceb525ccf938b7b6f5384b35373f1f46787467ebf5"
        imagePullPolicy: IfNotPresent
        command:
        - cilium-operator-generic
        args:
        - --config-dir=/tmp/cilium/config-map
        - --debug=$(CILIUM_DEBUG)
        env:
        - name: K8S_NODE_NAME
          valueFrom:
            fieldRef:
              apiVersion: v1
              fieldPath: spec.nodeName
        - name: CILIUM_K8S_NAMESPACE
          valueFrom:
            fieldRef:
              apiVersion: v1
              fieldPath: metadata.namespace
        - name: CILIUM_DEBUG
          valueFrom:
            configMapKeyRef:
              key: debug
              name: cilium-config
              optional: true
        livenessProbe:
          httpGet:
            host: "127.0.0.1"
            path: /healthz
            port: 9234
            scheme: HTTP
          initialDelaySeconds: 60
          periodSeconds: 10
          timeoutSeconds: 3
        volumeMounts:
        - name: cilium-config-path
          mountPath: /tmp/cilium/config-map
          readOnly: true
        terminationMessagePolicy: FallbackToLogsOnError
      hostNetwork: true
      restartPolicy: Always
      priorityClassName: system-cluster-critical
      serviceAccount: "cilium-operator"
      serviceAccountName: "cilium-operator"
      # In HA mode, cilium-operator pods must not be scheduled on the same
      # node as they will clash with each other.
      affinity:
        podAntiAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
          - labelSelector:
              matchLabels:
                io.cilium/app: operator
            topologyKey: kubernetes.io/hostname
      nodeSelector:
        kubernetes.io/os: linux
      tolerations:
        - operator: Exists
      volumes:
        # To read the configuration from the config map
      - name: cilium-config-path
        configMap:
          name: cilium-config
---
# Source: cilium/templates/hubble-relay/deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: hubble-relay
  namespace: kube-system
  labels:
    k8s-app: hubble-relay
    app.kubernetes.io/name: hubble-relay
    app.kubernetes.io/part-of: cilium
spec:
  replicas: 1
  selector:
    matchLabels:
      k8s-app: hubble-relay
  strategy:
    rollingUpdate:
      maxUnavailable: 1
    type: RollingUpdate
  template:
    metadata:
      annotations:
      labels:
        k8s-app: hubble-relay
        app.kubernetes.io/name: hubble-relay
        app.kubernetes.io/part-of: cilium
    spec:
      containers:
        - name: hubble-relay
          image: "quay.io/cilium/hubble-relay:v1.13.0@sha256:bc00f086285d2d287dd662a319d3dbe90e57179515ce8649425916aecaa9ac3c"
          imagePullPolicy: IfNotPresent
          command:
            - hubble-relay
          args:
            - serve
          ports:
            - name: grpc
              containerPort: 4245
          readinessProbe:
            tcpSocket:
              port: grpc
          livenessProbe:
            tcpSocket:
              port: grpc
          volumeMounts:
          - name: config
            mountPath: /etc/hubble-relay
            readOnly: true
          - name: tls
            mountPath: /var/lib/hubble-relay/tls
            readOnly: true
          terminationMessagePolicy: FallbackToLogsOnError
      restartPolicy: Always
      priorityClassName: 
      serviceAccount: "hubble-relay"
      serviceAccountName: "hubble-relay"
      automountServiceAccountToken: false
      terminationGracePeriodSeconds: 1
      affinity:
        podAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
          - labelSelector:
              matchLabels:
                k8s-app: cilium
            topologyKey: kubernetes.io/hostname
      nodeSelector:
        kubernetes.io/os: linux
      volumes:
      - name: config
        configMap:
          name: hubble-relay-config
          items:
          - key: config.yaml
            path: config.yaml
      - name: tls
        projected:
          # note: the leading zero means this number is in octal representation: do not remove it
          defaultMode: 0400
          sources:
          - secret:
              name: hubble-relay-client-certs
              items:
                - key: ca.crt
                  path: hubble-server-ca.crt
                - key: tls.crt
                  path: client.crt
                - key: tls.key
                  path: client.key
---
# Source: cilium/templates/hubble-ui/deployment.yaml
kind: Deployment
apiVersion: apps/v1
metadata:
  name: hubble-ui
  namespace: kube-system
  labels:
    k8s-app: hubble-ui
    app.kubernetes.io/name: hubble-ui
    app.kubernetes.io/part-of: cilium
spec:
  replicas: 1
  selector:
    matchLabels:
      k8s-app: hubble-ui
  template:
    metadata:
      annotations:
      labels:
        k8s-app: hubble-ui
        app.kubernetes.io/name: hubble-ui
        app.kubernetes.io/part-of: cilium
    spec:
      securityContext:
        fsGroup: 1001
        runAsGroup: 1001
        runAsUser: 1001
      priorityClassName: 
      serviceAccount: "hubble-ui"
      serviceAccountName: "hubble-ui"
      containers:
      - name: frontend
        image: "quay.io/cilium/hubble-ui:v0.10.0@sha256:118ad2fcfd07fabcae4dde35ec88d33564c9ca7abe520aa45b1eb13ba36c6e0a"
        imagePullPolicy: IfNotPresent
        ports:
        - name: http
          containerPort: 8081
        volumeMounts:
          - name: hubble-ui-nginx-conf
            mountPath: /etc/nginx/conf.d/default.conf
            subPath: nginx.conf
          - name: tmp-dir
            mountPath: /tmp
        terminationMessagePolicy: FallbackToLogsOnError
      - name: backend
        image: "quay.io/cilium/hubble-ui-backend:v0.10.0@sha256:cc5e2730b3be6f117b22176e25875f2308834ced7c3aa34fb598aa87a2c0a6a4"
        imagePullPolicy: IfNotPresent
        env:
        - name: EVENTS_SERVER_PORT
          value: "8090"
        - name: FLOWS_API_ADDR
          value: "hubble-relay:80"
        ports:
        - name: grpc
          containerPort: 8090
        volumeMounts:
        terminationMessagePolicy: FallbackToLogsOnError
      nodeSelector:
        kubernetes.io/os: linux
      volumes:
      - configMap:
          defaultMode: 420
          name: hubble-ui-nginx
        name: hubble-ui-nginx-conf
      - emptyDir: {}
        name: tmp-dir
---
# Source: none.. I made it up
kind: Service
apiVersion: v1
metadata:
  name: hubble-ui-node
  namespace: kube-system
  labels:
    k8s-app: hubble-ui
    app.kubernetes.io/name: hubble-ui
    app.kubernetes.io/part-of: cilium
spec:
  type: "NodePort"
  selector:
    k8s-app: hubble-ui
  ports:
    - name: http
      port: 8081
      nodePort: 30010`

var tetragon = `# Source: tetragon/templates/serviceaccount.yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: tetragon
  namespace: default
  labels:
    helm.sh/chart: tetragon-0.8.4
    app.kubernetes.io/name: tetragon
    app.kubernetes.io/instance: tetragon
    app.kubernetes.io/managed-by: Helm
---
# Source: tetragon/templates/tetragon_configmap.yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: tetragon-config
  namespace: default
  labels:
    helm.sh/chart: tetragon-0.8.4
    app.kubernetes.io/name: tetragon
    app.kubernetes.io/instance: tetragon
    app.kubernetes.io/managed-by: Helm
data:
  procfs: /procRoot
  enable-process-cred: "false"
  enable-process-ns: "false"
  process-cache-size: "65536"
  export-filename: /var/run/cilium/tetragon/tetragon.log
  export-file-max-size-mb: "10"
  export-file-max-backups: "5"
  export-file-compress: "false"
  export-allowlist: |-
    {"event_set":["PROCESS_EXEC", "PROCESS_EXIT", "PROCESS_KPROBE"]}
  export-denylist: |-
    {"health_check":true}
    {"namespace":["", "cilium", "kube-system"]}
  field-filters: |-
    {}
  export-rate-limit: "-1"
  enable-k8s-api: "true"
  metrics-server: :2112
  server-address: localhost:54321
  gops-address: localhost:8118
---
# Source: tetragon/templates/clusterrole.yaml
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: tetragon
  labels:
    helm.sh/chart: tetragon-0.8.4
    app.kubernetes.io/name: tetragon
    app.kubernetes.io/instance: tetragon
    app.kubernetes.io/managed-by: Helm
rules:
  - apiGroups:
      - ""
    resources:
      - pods
    verbs:
      - get
      - list
      - watch
  - apiGroups:
      - cilium.io
    resources:
      - tracingpolicies
    verbs:
      - get
      - list
      - watch
  # We need to split out the create permission and enforce it without resourceNames since
  # the name would not be known at resource creation time
  - apiGroups:
      - apiextensions.k8s.io
    resources:
      - customresourcedefinitions
    verbs:
      - create
  - apiGroups:
      - apiextensions.k8s.io
    resources:
      - customresourcedefinitions
    verbs:
      - create
  - apiGroups:
      - apiextensions.k8s.io
    resources:
      - customresourcedefinitions
    resourceNames:
      - tracingpolicies.cilium.io
    verbs:
      - update
      - get
      - list
---
# Source: tetragon/templates/clusterrolebinding.yml
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: tetragon
  labels:
    helm.sh/chart: tetragon-0.8.4
    app.kubernetes.io/name: tetragon
    app.kubernetes.io/instance: tetragon
    app.kubernetes.io/managed-by: Helm
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: tetragon
subjects:
  - kind: ServiceAccount
    namespace: default
    name: tetragon
---
# Source: tetragon/templates/service.yaml
apiVersion: v1
kind: Service
metadata:
  labels:
    helm.sh/chart: tetragon-0.8.4
    app.kubernetes.io/name: tetragon
    app.kubernetes.io/instance: tetragon
    app.kubernetes.io/managed-by: Helm
  name: tetragon
  namespace: default
spec:
  ports:
    - name: metrics
      port: 2112
      protocol: TCP
      targetPort: 2112
  selector:
    helm.sh/chart: tetragon-0.8.4
    app.kubernetes.io/name: tetragon
    app.kubernetes.io/instance: tetragon
    app.kubernetes.io/managed-by: Helm
  type: ClusterIP
---
# Source: tetragon/templates/daemonset.yaml
apiVersion: apps/v1
kind: DaemonSet
metadata:
  labels:
    helm.sh/chart: tetragon-0.8.4
    app.kubernetes.io/name: tetragon
    app.kubernetes.io/instance: tetragon
    app.kubernetes.io/managed-by: Helm
  name: tetragon
  namespace: default
spec:
  selector:
    matchLabels:
      app.kubernetes.io/name: tetragon
      app.kubernetes.io/instance: tetragon
  template:
    metadata:
      labels:
        helm.sh/chart: tetragon-0.8.4
        app.kubernetes.io/name: tetragon
        app.kubernetes.io/instance: tetragon
        app.kubernetes.io/managed-by: Helm
    spec:
      serviceAccountName: tetragon
      initContainers:
      - name: tetragon-operator
        command:
        - tetragon-operator
        image: "quay.io/cilium/tetragon-operator:v0.8.4"
        imagePullPolicy: IfNotPresent
      containers:
      - name: export-stdout
        image: "quay.io/cilium/hubble-export-stdout:v1.0.2"
        imagePullPolicy: IfNotPresent
        env:
          []
        securityContext:
          {}
        resources:
          {}
        command:
          - hubble-export-stdout
        args:
          - /var/run/cilium/tetragon/tetragon.log
        volumeMounts:
          - name: export-logs
            mountPath: /var/run/cilium/tetragon
      - name: tetragon
        securityContext:
          privileged: true
        image: "quay.io/cilium/tetragon:v0.8.4"
        imagePullPolicy: IfNotPresent
        command:
          - /usr/bin/tetragon
        args:
          - --config-dir=/etc/tetragon/tetragon.conf.d/
        volumeMounts:
          - mountPath: /var/lib/tetragon/metadata
            name: metadata-files
          - mountPath: /etc/tetragon/tetragon.conf.d/
            name: tetragon-config
            readOnly: true
          - mountPath: /sys/fs/bpf
            mountPropagation: Bidirectional
            name: bpf-maps
          - mountPath: "/var/run/cilium"
            name: cilium-run
          - mountPath: /var/run/cilium/tetragon
            name: export-logs
          - mountPath: "/procRoot"
            name: host-proc
        env:
          - name: NODE_NAME
            valueFrom:
              fieldRef:
                  fieldPath: spec.nodeName
        livenessProbe:
           exec:
             command:
             - tetra
             - status
             - --server-address
             - localhost:54321
      tolerations:
      - operator: Exists
      hostNetwork: true
      dnsPolicy: Default
      terminationGracePeriodSeconds: 1
      volumes:
      - name: cilium-run
        hostPath:
          path: /var/run/cilium
          type: DirectoryOrCreate
      - name: export-logs
        hostPath:
          path: /var/run/cilium/tetragon
          type: DirectoryOrCreate
      - name: tetragon-config
        configMap:
          name: tetragon-config
      - name: bpf-maps
        hostPath:
          path: /sys/fs/bpf
          type: DirectoryOrCreate
      - name: host-proc
        hostPath:
          path: /proc
          type: Directory
      - emptyDir: {}
        name: metadata-files`