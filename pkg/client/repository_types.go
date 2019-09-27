/*
 * Copyright (C) 2019  Rohith Jayawardene <gambol99@gmail.com>
 *
 * This program is free software; you can redistribute it and/or
 * modify it under the terms of the GNU General Public License
 * as published by the Free Software Foundation; either version 2
 * of the License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package client

import (
	"fmt"
	"time"
)

// Repository defines a repository
type Repository struct {
	CanAdmin       bool                      `json:"can_admin"`
	CanWrite       bool                      `json:"can_write"`
	Description    string                    `json:"description"`
	IsOrganization bool                      `json:"is_organization"`
	IsPublic       bool                      `json:"is_public"`
	IsStarred      bool                      `json:"is_starred"`
	Kind           string                    `json:"kind"`
	Name           string                    `json:"name"`
	Namespace      string                    `json:"namespace"`
	StatusToken    string                    `json:"status_token"`
	TagExpirationS int                       `json:"tag_expiration_s"`
	Tags           map[string]*RepositoryTag `json:"tags"`
	TrustEnabled   bool                      `json:"trust_enabled"`
}

// ImageAnalysis defines the results of an image scan
type ImageAnalysis struct {
	Status string `json:"status"`
	Data   struct {
		Layer struct {
			Features         []*ImageFeature `json:"Features"`
			IndexedByVersion int             `json:"IndexedByVersion"`
			Name             string          `json:"Name"`
			NamespaceName    string          `json:"NamespaceName"`
			ParentName       string          `json:"ParentName"`
		} `json:"Layer"`
	} `json:"data"`
}

// ImageFeature defined the feature in a image layer
type ImageFeature struct {
	AddedBy         string                  `json:"AddedBy"`
	Name            string                  `json:"Name"`
	NamespaceName   string                  `json:"NamespaceName"`
	Version         string                  `json:"Version"`
	VersionFormat   string                  `json:"VersionFormat"`
	Vulnerabilities []*ImageVulnerabilities `json:"Vulnerabilities,omitempty"`
}

// ImageVulnerabilities defines a vulnerability in the layer
type ImageVulnerabilities struct {
	NamespaceName string `json:"NamespaceName"`
	Link          string `json:"Link"`
	Name          string `json:"Name"`
	FixedBy       string `json:"FixedBy"`
	Severity      string `json:"Severity"`
}

// RepositoryTag defines a repository tag
type RepositoryTag struct {
	ImageID        string `json:"image_id"`
	Modified       string `json:"last_modified"`
	LastModified   time.Time
	Name           string `json:"name"`
	ManifestDigest string `json:"manifest_digest"`
	Size           int    `json:"size"`
}

// NewRepo defines a new repository
type NewRepo struct {
	RepoKind    string `json:"repo_kind"`
	Namespace   string `json:"namespace"`
	Visibility  string `json:"visibility"`
	Repository  string `json:"repository"`
	Description string `json:"description"`
}

// Name returns the full name of the repository
func (n *NewRepo) Name() string {
	return fmt.Sprintf("%s/%s", n.Namespace, n.Repository)
}

// Permission defines a user permission on a repository
type Permission struct {
	IsOrgMember bool   `json:"is_org_member"`
	Role        string `json:"role"`
	Name        string `json:"name"`
	IsRobot     bool   `json:"is_robot"`
}

// RepositoryList defined a list of repositories
type RepositoryList struct {
	NextPage     string `json:"next_page"`
	Repositories []struct {
		Kind        string `json:"kind"`
		Name        string `json:"name"`
		Namespace   string `json:"namespace"`
		IsPublic    bool   `json:"is_public"`
		IsStarred   bool   `json:"is_starred"`
		Description string `json:"description"`
	} `json:"repositories"`
}
