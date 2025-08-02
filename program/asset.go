package program

import (
	"encoding/json"
	"fmt"
)

// AssetMetadata is an interface that all asset metadata types must implement
type AssetMetadata interface {
	Type() AssetType
}

type Asset struct {
	Type     AssetType     `json:"type"`
	Metadata AssetMetadata `json:"metadata"`
}

type AssetType int

const (
	File AssetType = iota
	DockerCompose
	LibvirtQemu
	LibvirtLXC
)

var AssetTypeNames = map[AssetType]string{
	File:          "file",
	DockerCompose: "docker-compose",
	LibvirtQemu:   "libvirt-qemu",
	LibvirtLXC:    "libvirt-lxc",
}

func (t AssetType) toString() (string, error) {
	if t < 0 || int(t) >= len(AssetTypeNames) {
		return "", fmt.Errorf("invalid asset type: %d", t)
	}
	return AssetTypeNames[t], nil
}

func toAssetType(s string) (AssetType, error) {
	for t, name := range AssetTypeNames {
		if name == s {
			return t, nil
		}
	}
	return -1, fmt.Errorf("invalid asset type: %s", s)
}

type FileAssetMetadata struct {
	Path string `json:"path"`
}

func (f FileAssetMetadata) Type() AssetType {
	return File
}

type DockerComposeAssetMetadata struct {
	Path string `json:"path"`
}

func (d DockerComposeAssetMetadata) Type() AssetType {
	return DockerCompose
}

type LibvirtQemuAssetMetadata struct {
	Path string `json:"path"`
}

func (l LibvirtQemuAssetMetadata) Type() AssetType {
	return LibvirtQemu
}

type LibvirtLXCAssetMetadata struct {
	Path string `json:"path"`
}

func (l LibvirtLXCAssetMetadata) Type() AssetType {
	return LibvirtLXC
}

func (a *Asset) MarshalJSON() ([]byte, error) {
	type MarshalledAsset struct {
		Type     string      `json:"type"`
		Metadata interface{} `json:"metadata"`
	}

	typeStr, err := a.Type.toString()
	if err != nil {
		return nil, err
	}

	temp := MarshalledAsset{
		Type:     typeStr,
		Metadata: a.Metadata,
	}

	return json.Marshal(temp)
}

func (a *Asset) UnmarshalJSON(data []byte) error {
	type UnmarshalledAsset struct {
		Type     string          `json:"type"`
		Metadata json.RawMessage `json:"metadata"`
	}

	var temp UnmarshalledAsset
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	assetType, err := toAssetType(temp.Type)
	if err != nil {
		return err
	}

	a.Type = assetType

	switch assetType {
	case File:
		var metadata FileAssetMetadata
		if err := json.Unmarshal(temp.Metadata, &metadata); err != nil {
			return fmt.Errorf("failed to unmarshal file metadata: %w", err)
		}
		a.Metadata = metadata
	case DockerCompose:
		var metadata DockerComposeAssetMetadata
		if err := json.Unmarshal(temp.Metadata, &metadata); err != nil {
			return fmt.Errorf("failed to unmarshal docker-compose metadata: %w", err)
		}
		a.Metadata = metadata
	case LibvirtQemu:
		var metadata LibvirtQemuAssetMetadata
		if err := json.Unmarshal(temp.Metadata, &metadata); err != nil {
			return fmt.Errorf("failed to unmarshal libvirt-qemu metadata: %w", err)
		}
		a.Metadata = metadata
	case LibvirtLXC:
		var metadata LibvirtLXCAssetMetadata
		if err := json.Unmarshal(temp.Metadata, &metadata); err != nil {
			return fmt.Errorf("failed to unmarshal libvirt-lxc metadata: %w", err)
		}
		a.Metadata = metadata
	default:
		return fmt.Errorf("unimplemented asset type unmarshaller: %s", temp.Type)
	}

	return nil
}
