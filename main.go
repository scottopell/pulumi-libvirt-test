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

	domainArgs := &libvirt.DomainArgs{
		Name:     pulumi.String("HelloUbuntu"),
		Arch:     pulumi.String("aarch64"),
		Emulator: pulumi.String("/opt/homebrew/bin/qemu-system-aarch64"),
	}

	// TODO play with Machine value here, this is emitting
	_, err = libvirt.NewDomain(ctx, "ubuntu", domainArgs, pulumi.Provider(provider))
	if err != nil {
		return err
	}

	fmt.Println("Done.")

	return nil
}

func main() {
	pulumi.Run(providerAndDomain)
}
