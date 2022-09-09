package main

import (
	"fmt"

	"github.com/pulumi/pulumi-libvirt/sdk/go/libvirt"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func providerAndDomain(ctx *pulumi.Context) error {
	provider, err := libvirt.NewProvider(ctx, "provider", &libvirt.ProviderArgs{Uri: pulumi.String("qemu:///session?socket=/opt/homebrew/var/run/libvirt/libvirt-sock")})
	if err != nil {
		return err
	}

	// TODO play with Machine DomainArg here, this is emitting kvm by default
	_, err = libvirt.NewDomain(ctx, "ubuntu", &libvirt.DomainArgs{Name: pulumi.String("HelloUbuntu")}, pulumi.Provider(provider))
	if err != nil {
		return err
	}

	fmt.Println("Done.")

	return nil
}

func main() {
	pulumi.Run(providerAndDomain)
}
