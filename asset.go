package program

import (
	"encoding/json"
	"fmt"
)

type Asset struct {
	Type     AssetType              `json:"type"`
	Metadata map[string]interface{} `json:"metadata"`
}

type AssetType int

const (
	File AssetType = iota
	DockerCompose
	LibvirtQemu
	LibvirtKvm
)

var AssetTypeNames = map[AssetType]string{
	File:          "file",
	DockerCompose: "docker-compose",
	LibvirtQemu:   "libvirt-qemu",
	LibvirtKvm:    "libvirt-kvm",
}

func (t AssetType) String() string {
	return AssetTypeNames[t]
}

type FileAssetMetadata struct {
	Path string `json:"path"`
}

type DockerComposeAssetMetadata struct {
	Path string `json:"path"`
}

type LibvirtQemuAssetMetadata struct {
	Path string `json:"path"`
}

type LibvirtKvmAssetMetadata struct {
	Path string `json:"path"`
}

// MarshalJSON implements custom JSON marshaling for Asset
func (a *Asset) MarshalJSON() ([]byte, error) {
	// Create a temporary struct for marshaling
	type assetAlias Asset

	// Validate the asset type
	if _, exists := AssetTypeNames[a.Type]; !exists {
		return nil, fmt.Errorf("invalid asset type: %d", a.Type)
	}

	// Validate metadata based on type
	if err := a.validateMetadata(); err != nil {
		return nil, fmt.Errorf("invalid metadata for type %s: %w", a.Type.String(), err)
	}

	return json.Marshal((*assetAlias)(a))
}

// UnmarshalJSON implements custom JSON unmarshaling for Asset
func (a *Asset) UnmarshalJSON(data []byte) error {
	// Create a temporary struct for unmarshaling
	type assetAlias Asset

	var temp assetAlias
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	// Validate the asset type
	if _, exists := AssetTypeNames[temp.Type]; !exists {
		return fmt.Errorf("invalid asset type: %d", temp.Type)
	}

	// Convert metadata to proper struct based on type
	asset := Asset(temp)
	if err := asset.convertMetadata(); err != nil {
		return fmt.Errorf("failed to convert metadata for type %s: %w", asset.Type.String(), err)
	}

	*a = asset
	return nil
}

// validateMetadata validates that the metadata matches the expected structure for the asset type
func (a *Asset) validateMetadata() error {
	switch a.Type {
	case File:
		if _, ok := a.Metadata["path"]; !ok {
			return fmt.Errorf("file asset requires 'path' field in metadata")
		}
		if path, ok := a.Metadata["path"].(string); !ok || path == "" {
			return fmt.Errorf("file asset requires non-empty string 'path' field")
		}

	case DockerCompose:
		if _, ok := a.Metadata["path"]; !ok {
			return fmt.Errorf("docker-compose asset requires 'path' field in metadata")
		}
		if path, ok := a.Metadata["path"].(string); !ok || path == "" {
			return fmt.Errorf("docker-compose asset requires non-empty string 'path' field")
		}

	case LibvirtQemu:
		if _, ok := a.Metadata["path"]; !ok {
			return fmt.Errorf("libvirt-qemu asset requires 'path' field in metadata")
		}
		if path, ok := a.Metadata["path"].(string); !ok || path == "" {
			return fmt.Errorf("libvirt-qemu asset requires non-empty string 'path' field")
		}

	case LibvirtKvm:
		if _, ok := a.Metadata["path"]; !ok {
			return fmt.Errorf("libvirt-kvm asset requires 'path' field in metadata")
		}
		if path, ok := a.Metadata["path"].(string); !ok || path == "" {
			return fmt.Errorf("libvirt-kvm asset requires non-empty string 'path' field")
		}

	default:
		return fmt.Errorf("unknown asset type: %d", a.Type)
	}

	return nil
}

// convertMetadata converts the generic metadata map to the appropriate typed struct
func (a *Asset) convertMetadata() error {
	switch a.Type {
	case File:
		if path, ok := a.Metadata["path"].(string); ok {
			// Metadata is already in the correct format
			_ = path // Use path to avoid unused variable warning
		} else {
			return fmt.Errorf("invalid file metadata format")
		}

	case DockerCompose:
		if path, ok := a.Metadata["path"].(string); ok {
			// Metadata is already in the correct format
			_ = path // Use path to avoid unused variable warning
		} else {
			return fmt.Errorf("invalid docker-compose metadata format")
		}

	case LibvirtQemu:
		if path, ok := a.Metadata["path"].(string); ok {
			// Metadata is already in the correct format
			_ = path // Use path to avoid unused variable warning
		} else {
			return fmt.Errorf("invalid libvirt-qemu metadata format")
		}

	case LibvirtKvm:
		if path, ok := a.Metadata["path"].(string); ok {
			// Metadata is already in the correct format
			_ = path // Use path to avoid unused variable warning
		} else {
			return fmt.Errorf("invalid libvirt-kvm metadata format")
		}

	default:
		return fmt.Errorf("unknown asset type: %d", a.Type)
	}

	return nil
}

// GetFileMetadata returns the metadata as FileAssetMetadata if the type is File
func (a *Asset) GetFileMetadata() (*FileAssetMetadata, error) {
	if a.Type != File {
		return nil, fmt.Errorf("asset is not of type File")
	}

	if err := a.validateMetadata(); err != nil {
		return nil, err
	}

	path, _ := a.Metadata["path"].(string)
	return &FileAssetMetadata{Path: path}, nil
}

// GetDockerComposeMetadata returns the metadata as DockerComposeAssetMetadata if the type is DockerCompose
func (a *Asset) GetDockerComposeMetadata() (*DockerComposeAssetMetadata, error) {
	if a.Type != DockerCompose {
		return nil, fmt.Errorf("asset is not of type DockerCompose")
	}

	if err := a.validateMetadata(); err != nil {
		return nil, err
	}

	path, _ := a.Metadata["path"].(string)
	return &DockerComposeAssetMetadata{Path: path}, nil
}

// GetLibvirtQemuMetadata returns the metadata as LibvirtQemuAssetMetadata if the type is LibvirtQemu
func (a *Asset) GetLibvirtQemuMetadata() (*LibvirtQemuAssetMetadata, error) {
	if a.Type != LibvirtQemu {
		return nil, fmt.Errorf("asset is not of type LibvirtQemu")
	}

	if err := a.validateMetadata(); err != nil {
		return nil, err
	}

	path, _ := a.Metadata["path"].(string)
	return &LibvirtQemuAssetMetadata{Path: path}, nil
}

// GetLibvirtKvmMetadata returns the metadata as LibvirtKvmAssetMetadata if the type is LibvirtKvm
func (a *Asset) GetLibvirtKvmMetadata() (*LibvirtKvmAssetMetadata, error) {
	if a.Type != LibvirtKvm {
		return nil, fmt.Errorf("asset is not of type LibvirtKvm")
	}

	if err := a.validateMetadata(); err != nil {
		return nil, err
	}

	path, _ := a.Metadata["path"].(string)
	return &LibvirtKvmAssetMetadata{Path: path}, nil
}
