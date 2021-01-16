package main

import (
	"encoding/json"
	"fmt"
)

type TestByte struct {
	KubeConfig  string `json:"kube_config"`
	ClusterName string `json:"cluster_name"`
}

func main() {
	var cluster = `apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUN5RENDQWJDZ0F3SUJBZ0lCQURBTkJna3Foa2lHOXcwQkFRc0ZBREFWTVJNd0VRWURWUVFERXdwcmRXSmwKY201bGRHVnpNQjRYRFRJd01UQXpNVEV6TXpReU0xb1hEVE13TVRBeU9URXpNelF5TTFvd0ZURVRNQkVHQTFVRQpBeE1LYTNWaVpYSnVaWFJsY3pDQ0FTSXdEUVlKS29aSWh2Y05BUUVCQlFBRGdnRVBBRENDQVFvQ2dnRUJBTUlXCnVqcU1vZGdvU2k5eVovSkg2SlJjWFBKbk1yRkd1V2tMb1EzMk9NazdLWmRtVHFnTjBZOGlOTVg4OGZveFV3dkkKUkEyNC9tMnc4dWFYQTErVUwwZkRGWHZvTjhWTFl2dGJiK3JncjRzZ1JvRXQwQjErNW4rdDBLT2VLZTVXZ1ZqSAp2WmJKcjNTTFBmRUQ1ek5LbEwyMEJlQ2RYS1duNTZEMkg1b2dveFZnU3lUS2VodituK2FLVDFSQ1QwdzN4aFRXCkFsd0xtd2loeHdOWkN3bFI4bGt4M1hkOGYxTCtsSXVBRjNMUCsvUndqdDJ6NXk2TFcwdzZ6ZW5HTGpWM2pUcHAKZlM1UlBEcG5pZDhUNmVzMlBJWDNFOWVLeVFnQ2ZFRUM0SUp5WFR1dXgvcnI4WkNqVmFEUUFucndmdWh3U1huWAo5cU5ydWlCYS9NWnpHTDJiSG5FQ0F3RUFBYU1qTUNFd0RnWURWUjBQQVFIL0JBUURBZ0trTUE4R0ExVWRFd0VCCi93UUZNQU1CQWY4d0RRWUpLb1pJaHZjTkFRRUxCUUFEZ2dFQkFMbWhjdTFZOW9lZGdIUzQxdTlWTDZleHd4d0UKdDN6OGoxQzBMa0o0R1czU2ljSzdqZ1FlcVR6VEZHaWxLelVQMXdVcDFwWVhGK21Nd1QzZmMwMXRaZWROS3diUwpYL1drNjd4VTBXbi9ucVFrSmU3NmMxd1lBSEwzNHQ3ZGhwbXRtRTdiaEhnVGx5Tkl2YkhRaXhqZFpBQTk0RG9MCm5xWW0vUDJtUVdHNDgrUmRkQlJLY3VtVDd2WXljU2VFVUg1UkFBTXA4cmRUZTFlNXd4OEhrTWltRDhTVXJsZGsKYkZ6V2pZOEpnNjRmTTJRTGJHNHFuN2ZTUTcrOHhaVkh4MHlzZ2VJUGx0ZGg1V1QwOWtLYWV4TVhhNDVqQ0xGcwpXU3E4aE5aSzl5T08wZE1WcUtLa1FJNHRKN1MrRlJUc09mZmFqdURyQWVXWWwrdXJ1UDRNYlJYQ0l2bz0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=
    server: https://192.168.56.5:6443
  name: kubernetes
contexts:
- context:
    cluster: kubernetes
    user: kubernetes-admin
  name: kubernetes-admin@kubernetes
current-context: kubernetes-admin@kubernetes
kind: Config
preferences: {}
users:
- name: kubernetes-admin
  user:
    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUM4akNDQWRxZ0F3SUJBZ0lJSnN6Wlh1NGhVdE13RFFZSktvWklodmNOQVFFTEJRQXdGVEVUTUJFR0ExVUUKQXhNS2EzVmlaWEp1WlhSbGN6QWVGdzB5TURFd016RXhNek0wTWpOYUZ3MHlNVEV3TXpFeE16TTBNalZhTURReApGekFWQmdOVkJBb1REbk41YzNSbGJUcHRZWE4wWlhKek1Sa3dGd1lEVlFRREV4QnJkV0psY201bGRHVnpMV0ZrCmJXbHVNSUlCSWpBTkJna3Foa2lHOXcwQkFRRUZBQU9DQVE4QU1JSUJDZ0tDQVFFQXhzQmpvVkQvcEgvUXY0R3AKV016TnpyVEVYTk81Q0YvOGw2bFdvWjJpb2VIZlg1OEYyd3o1VlFSTlVkeXNSY1NmSEl6TDUwTGNxRzJDamFvMwpnbmp2K29GbUd1N01uaVNpcENHR1JCejFTYmN2Z2F1K1RYYXU2NE5CWjZnTmRYOTdoK1RneTJsSkRvdFVsYjh1Cm9EWmIrYUxwUklNL1B6S0J5NUp4TjdRTUNOQmExQjM5NWhiOGZYb1ZXL3lqY1dna1NpRUdDdHNzK0Q0aG1yUysKaWQ2bXdZaXFVRWRURzMxWkxrUE5JZ29RMnBPZFNqcmt1RlgxRy9tTVRadHhUNzh1ZGQ2UTF4UUZYMGtwRk5zNQpoenFYb3V5WHZmUVd5OWw1OWdtVlBzWWxuUEdTaDJrTThyM1o4a2lqSzZBMC9QN0VlekxaaTd5NXlFaWlUWmhmClhOWDI3UUlEQVFBQm95Y3dKVEFPQmdOVkhROEJBZjhFQkFNQ0JhQXdFd1lEVlIwbEJBd3dDZ1lJS3dZQkJRVUgKQXdJd0RRWUpLb1pJaHZjTkFRRUxCUUFEZ2dFQkFCMzhFVWd0eldBbnpEKzIzckRtOW9uWk1zSC9ucTJmTXNkWAo3djArMHhCRHYvWmY5RjdFb29JUHh3Si9mMG80bFdIMEZvMEZzYndObi9Zb2xib2tJZmJKT1pNVUVaS1Y0WktqCkFieE1ONTh4Snk3eUdFelQwYWw1T0IyVkVLTXBHUDZqOXJpbFlLcUdCWEpxRHZKak9CM045L2VMaEJxaERVTkoKME1kaHBDcHVEdkxwR1ZJTUd3R0t4VlduUTRRQUtmN1l2bW9lb3gvNnUyZ3lLMzZqb0xyZkVheUtUWFQyVmZEVgo0TEt4YmJrcVFaczJTV0tDRDF5VmxPUnRCUTJ2b3BETm4zWHNaOWYyMjcwT2MwWWttZ3VLVjNJclNmSkxFeWpnCnZWcDdYY3I5NTk0QjNGalRTWmFzSVhNY1drT08xaEx5Z3AxVDAzYytjeEo0UG5EVUhCST0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=
    client-key-data: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFb3dJQkFBS0NBUUVBeHNCam9WRC9wSC9RdjRHcFdNek56clRFWE5PNUNGLzhsNmxXb1oyaW9lSGZYNThGCjJ3ejVWUVJOVWR5c1JjU2ZISXpMNTBMY3FHMkNqYW8zZ25qditvRm1HdTdNbmlTaXBDR0dSQnoxU2JjdmdhdSsKVFhhdTY0TkJaNmdOZFg5N2grVGd5MmxKRG90VWxiOHVvRFpiK2FMcFJJTS9QektCeTVKeE43UU1DTkJhMUIzOQo1aGI4ZlhvVlcveWpjV2drU2lFR0N0c3MrRDRobXJTK2lkNm13WWlxVUVkVEczMVpMa1BOSWdvUTJwT2RTanJrCnVGWDFHL21NVFp0eFQ3OHVkZDZRMXhRRlgwa3BGTnM1aHpxWG91eVh2ZlFXeTlsNTlnbVZQc1lsblBHU2gya00KOHIzWjhraWpLNkEwL1A3RWV6TFppN3k1eUVpaVRaaGZYTlgyN1FJREFRQUJBb0lCQUVHb09NR0gwb3Z6TlhDdQpVdFZsQXZBd1ExWXZFMGN5L216VG9pek51SFlsK2ZiS3Y5SjNYTDY3SnhmT0FKTWp1elJoUnRMbEFhUWFXVENYCmRMTURaTE9xbjZKeUZPS0JVOEZJMmVRU2RPMmZPSEJyM2ZnVkp2L2lZbTd0MDFSbmd0V0h4TU10eXRhbExKSXoKa0NHWkU1VktQaWxGK2xUVVNKQ0psaUZDdDE4SmNBWHdyTlQ5Yzl4ZHZaRUhpL1QxSUhISy9icWpPN1hEZElSUgpvMjZaa0JvMUtOVEx0MEhyMGN3NEhmRHpYSnlwV29RaDludmFWRFZpNzBtc1hhZ0dCTUgrZTFlbmo0ZU01cWdoCnQ5blh4Y0gvTnNLclVIeWI0c3hUWVAvY04zR0pSNGlPUHdrV3dyQlQ5LzVwVGhSOFZINDIvYkY5SDczTWZRN2oKQWxCYzU3a0NnWUVBOVl4ZzcxVVpCZjZsbkJ3ZldVbWhTUW9pWWVieTZVWXgyMDdWeUZWVHNONlBxYU8ybFE0eApLckNoWFUrc0RQWHNGU2gxTUtrdDFEZWRLUURZblB1YUplZ0tYNWFIa0tKUmZFOWlEdGYzU09pV0hsNzAwOXphCll6UUZRZlkxSzUvS2U5SFl2ZXRKVy9OOTFIbnpBZzlMRWtKOU9nZFYwOWd6S3I5cENjMVlkbWNDZ1lFQXp6WVcKanlCSEdWc3NDK2xGdkVlQU9WeXA5YnZMRGNtbzA5QitFeXdpdXpGT2gxVHdoU0lZanlyUkhqSlJ2eDhwa0QvbApPRnFLUDA3V3NUMVlsU0pSOVFzaFM3ekVGb2Y5SEpoWXZvVDdwK2tzeGFQdTFkMlFFR1RRY281dlF1Ykx2bWZ0ClFRTmR2cGRrSkxsM0liRGZPT3RvT0oxcHNxZnRtNzdtcFpLR3k0c0NnWUJVaVg3NkJ2ODFmQ1ZDS01CMk9nNUkKR2pBYUtxdm91aDBDRnhNdEJJdHFza1VkcjFwMTJNdittWThFbENCMDJkbnNzdUFkMzdFQ2hoOGFkY3NkeU5XYQpSVmdFN3NzL2RWSytqSW9DK0VHUG5Wdkptd0dKUzFqNG5IbE82MWdFNjkrSFg0alNZZ1lBblVUdCtRbzc1RTRCClBGam11STNKQUg4bE1YOFZPOWcvOHdLQmdRQ3ZWMFMxcm9tYmtrUHVmR0lQVmEzU1hIQlJ5YXlrcVMyY3Q0UEgKODRRY1JUZi9Xc2dBMEg1WnY5Q2ZMZE53MkF6U3pNN0o1Z3V3VUtWcDY5Qm11bGtCODM0a1djTFFraGlFTndKNQpqWXhxWk8wRy9aYWdFTkhVbllxcU5ZOCszSkF5T2oxUTAybGpXelpad1R0bkhnVjJZL1dZV3RQVFl6aTBuMUVMCmNMZHh5d0tCZ0hWbDhTditiSlRLUTRlRDh6MzZsOFd2Ty96ckRhU1IvSkh3T3Z5czBEU0kzWkx1eGFsWEFOdUQKMWhwQTk4WGhzRW5URFhub3VBclNGcFk5bzNWeHdXaWRLUUVhNFBGbmRUWEtEdGhIU0JVRTB3UnhhWWkvOFcrWQpPUG1KWjRkSW1NT1JMWkRuOXEyeGp2S05NUC9nKzZNWGxDV1MxRGdVdTQ1ZmpmSndXSUhGCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==`
	var test = TestByte{
		KubeConfig: cluster,
	}
	data, err := json.Marshal(test)
	if err != nil {
		panic(err.Error())
	}
	fmt.Errorf("Test Success with Output: ", data)
}
