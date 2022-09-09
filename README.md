# Pulumi Libvirt Attempt

Pulumi lets you auto-configure some infrastructure and then interact with it.

I'm interested in accomplishing this with local resources compared to cloud
resources (which is the norm).

There is a [`libvirt` provider](https://github.com/pulumi/pulumi-libvirt) that
allows this to work using `libvirt` which itself is a wrapper around various
virtualization platforms.

## Getting an Ubuntu VM Running
In `main.go` is my best attempt so far at running a simple ubuntu VM on top of
qemu on macos (arm64).

### Macos stuff
- `brew install libvirt qemu`
- `brew services libvirt start`

The `qemu` session string is currently hardcoded to the macos location of the
socket.

## Status
Currently this code fails because its generating xml for `kvm`, not `qemu` (or
ideally `hvf`?)

Ran with `pulumi up -d -v=3 --tracing $PWD/out.log`

```
    debug: Generated XML for libvirt domain:
    debug:   <domain type="kvm">
    debug:       <name>HelloUbuntu</name>
    debug:       <memory unit="MiB">512</memory>
    debug:       <vcpu>1</vcpu>
    debug:       <os firmware="efi">
    debug:           <type>hvm</type>
    debug:       </os>
    debug:       <features>
    debug:           <pae></pae>
    debug:           <acpi></acpi>
    debug:           <apic></apic>
    debug:       </features>
    debug:       <cpu></cpu>
    debug:       <devices>
    debug:           <channel type="unix">
    debug:               <target type="virtio" name="org.qemu.guest_agent.0"></target>
    debug:           </channel>
    debug:           <rng model="virtio">
    debug:               <backend model="random">/dev/urandom</backend>
    debug:           </rng>
    debug:       </devices>
    debug:   </domain>
```
