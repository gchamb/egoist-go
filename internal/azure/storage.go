package azure

import (
	"fmt"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/container"
)

var AZURE_STORAGE_SERVICE_URL = fmt.Sprintf("https://%s.blob.core.windows.net", os.Getenv("AZURE_STORAGE_ACCOUNT_NAME"))
const (
	PROGRESS_ENTRY_CONTAINER = "progress-entry"
)

func GetProgressEntryContainer() (*container.Client, error) {
	creds, err := GenerateAzureCreds()
	if err != nil {
		return nil, err
	}
	return container.NewClient(fmt.Sprintf("%s/%s", AZURE_STORAGE_SERVICE_URL, PROGRESS_ENTRY_CONTAINER), creds, &container.ClientOptions{})
}

func GetBlobClient(containerName string, blobName string, isSaS bool) (*blob.Client, error) {
	creds, err := GenerateAzureCreds()
	if err != nil {
		return nil, err
	}
	blobUrl := fmt.Sprintf("%s/%s/%s", AZURE_STORAGE_SERVICE_URL, containerName,blobName)
	if isSaS {
		delCreds, err := blob.NewSharedKeyCredential(os.Getenv("AZURE_STORAGE_ACCOUNT_NAME"), os.Getenv("AZURE_ACCOUNT_KEY"))
		if err != nil {
			return nil, err
		}

		return blob.NewClientWithSharedKeyCredential(blobUrl, delCreds, &blob.ClientOptions{})
	}
	
	return blob.NewClient(blobUrl, creds, &blob.ClientOptions{})
}