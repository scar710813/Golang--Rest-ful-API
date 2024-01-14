package azure

import (
    "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2022-01-01/compute"
    "github.com/Azure/go-autorest/autorest/azure/auth"
)

func ListVMs(subscriptionID string) ([]compute.VirtualMachine, error) {
    authorizer, err := auth.NewAuthorizerFromEnvironment()
    if err != nil {
        return nil, err
    }

    vmClient := compute.NewVirtualMachinesClient(subscriptionID)
    vmClient.Authorizer = authorizer

    result, err := vmClient.ListAll(context.Background())
    if err != nil {
        return nil, err
    }

    return result.Values(), nil
}
