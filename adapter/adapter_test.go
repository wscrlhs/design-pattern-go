package adapter_test

import (
	"design-pattern/adapter"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestMac(t *testing.T) {
	mac := &adapter.Mac{}

	client := adapter.Client{}
	result := client.InsertLightningConnectorIntoComputer(mac)
	assert.Equal(t, "mac lightning", result)
}

func TestWindows(t *testing.T) {
	windows := &adapter.Windows{}
	windowsAdapter := adapter.WindowsAdapter{
		WindowMachine: windows,
	}

	client := adapter.Client{}
	result := client.InsertLightningConnectorIntoComputer(&windowsAdapter)
	assert.Equal(t, "windows lightning", result)
}

func Example() {

	client := &adapter.Client{}
	mac := &adapter.Mac{}

	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := &adapter.Windows{}
	windowsMachineAdapter := &adapter.WindowsAdapter{
		WindowMachine: windowsMachine,
	}

	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)

	// Output:
	// Client inserts Lightning connector into computer.
	// Lightning connector is plugged into mac machine.
	// Client inserts Lightning connector into computer.
	// Adapter converts Lightning signal to USB.
	// USB connector is plugged into windows machine.
}
