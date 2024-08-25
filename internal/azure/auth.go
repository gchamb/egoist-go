package azure

import (
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

func GenerateAzureCreds() (*azidentity.ClientSecretCredential, error) {
	return azidentity.NewClientSecretCredential(os.Getenv("AZURE_TENANT_ID"), os.Getenv("AZURE_CLIENT_ID"), os.Getenv("AZURE_CLIENT_SECRET"), &azidentity.ClientSecretCredentialOptions{})
}



