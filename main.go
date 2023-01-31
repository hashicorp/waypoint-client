package main

import (
	"log"
	"os"

	apiclient "github.com/hashicorp/waypoint-client/client"
	"github.com/hashicorp/waypoint-client/gen/client/waypoint"

	"golang.org/x/oauth2/clientcredentials"
)

//Testing generated API. Not a permanent file.

func main() {

	clientConfig := apiclient.WaypointClientConfig{
		HCPOrgId:     os.Getenv("HCP_ORG_ID"),
		HCPProjectId: os.Getenv("HCP_PROJECT_ID"),
		ClientCredentialsConfig: clientcredentials.Config{
			ClientID:     os.Getenv("SP_CLIENT_ID"),
			ClientSecret: os.Getenv("SP_CLIENT_SECRET"),
		},
		ServerUrl:         os.Getenv("SERVER_URL"),
		WaypointUserToken: os.Getenv("WAYPOINT_USER_TOKEN"),
	}

	client, err := apiclient.New(clientConfig)
	if err != nil {
		log.Fatal(err)
	}
	client.WaypointListProjects(waypoint.NewWaypointListProjectsParams())

	//transport := httptransport.New(apiclient.DefaultHost, apiclient.DefaultBasePath, apiclient.DefaultSchemes)
	//bearerTokenAuth := httptransport.BearerToken("eyJhbGciOiJSUzI1NiIsImtpZCI6InB1YmxpYzpjNGNjZGUxYy00YjMzLTRiMzEtOTJmMS01NDhmNjdiZGUyY2QiLCJ0eXAiOiJKV1QifQ.eyJhdWQiOlsiaHR0cHM6Ly9hcGkuaGFzaGljb3JwLmNsb3VkIiwiaHR0cHM6Ly9hdXRoLmhhc2hpY29ycC5jbG91ZCJdLCJjbGllbnRfaWQiOiIyMWQ4NjI2Mi02ZjE0LTRhMzAtYTkwZi0wN2UzZmRlOGIyM2QiLCJleHAiOjE2NzM5OTIwNTEsImV4dCI6eyJodHRwczovL2F1dGguaGFzaGljb3JwLmNvbS9hdXRoMC9jb25uZWN0aW9uL2lkIjoiY29uX2RMZnpLaEFhSWtjdHJIVFIiLCJodHRwczovL2F1dGguaGFzaGljb3JwLmNvbS9hdXRoMC9jb25uZWN0aW9uL25hbWUiOiJnaXRodWIiLCJodHRwczovL2F1dGguaGFzaGljb3JwLmNvbS9hdXRoMC9jb25uZWN0aW9uL3N0cmF0ZWd5IjoiZ2l0aHViIiwiaHR0cHM6Ly9hdXRoLmhhc2hpY29ycC5jb20vaWRwIjoiYXV0aDAiLCJodHRwczovL2F1dGguaGFzaGljb3JwLmNvbS9zdWIiOiJnaXRodWJ8MTY5NTQ3OCJ9LCJpYXQiOjE2NzM5ODg0NTAsImlzcyI6Imh0dHBzOi8vYXV0aC5pZHAuaGFzaGljb3JwLmNvbS8iLCJqdGkiOiIxMGVmYmQ0NC01ZGQ0LTRiNTUtOTZkYy01MjA1NDNiNTBlOGQiLCJuYmYiOjE2NzM5ODg0NTAsInNjcCI6WyJvcGVuaWQiLCJvZmZsaW5lIl0sInN1YiI6IjQ4YWRkODRmLTEwMWYtNDcyNC1iMzM5LTVmNDZmOWZhMzYzYSJ9.PSWId1QIoB1maNbaTT93BjAFZ8-rtVjNTOKBCivTjh1w6eapv8lgcCV7QTh5S-GchzkhjcH778B9r5rxoQufFRfLBGb6RWBW7PzbKv9oAN88zIQrqjs3HHIujqCwk4zqXtMLva7NHAWFjpp4igtFkjNTWu0G2uTp_Y1RAj-vXkFgmWTHUubByJBPlIfdAr0ikr461W-1NpwmGa4tMLGam-f5QPDuiauCceivca8pklteNE31OMGjLMFXfLPvRpknyYevMPYQSlFi8gbk12yx-ngCguqzwonuO1HRpPZZjNeDGkGGcK0UtNbFogDMojx05BhGBawiw_36OcZdDNJg-jmE4edynu7QS5uwWcNVltnw_twhPPWr2REFImJVzaHwM2PLDeLzSMNabak1xeiK3DYc3ydbyk7GsXaSZVnmpAie9h--h2oK7FwZB0MDF7HOM7hC48FBxXFKeaO1cNeQ4Ix8HuFggFtJAYvNEBj6vHqIH7Dow9mPkpn9wvVwgvFbNDm4FRO7JezZOOb1blmozFnyNcZoPm2jYBU0XaFbmj_3lhL24ThHR2jCw8cuGYAxRt9Rx-Ai08vRp0quSSK6l-cFymV9TA8uq9sRgJeWk9Rfkq2Y0ztx96MP2GIzB0pDGANO6MGeI3wNHS99PyyVszJEiVQVH54r20MiIY3j-jg")
	//client := apiclient.New(transport, strfmt.Default)
	//
	//resp, err := client.Waypoint.WaypointListProjects(waypoint.NewWaypointListProjectsParams().WithNamespaceID("54c86519-3257-4d31-9006-91f9e003a1fd"), bearerTokenAuth)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Printf("%#v\n", resp.Payload)

}
