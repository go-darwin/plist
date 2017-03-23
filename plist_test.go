// Copyright 2017 The go-darwin Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plist

import (
	"testing"
)

var testPlist = []byte(`<plist version="1.0">
    <dict>
        <key>BucketUUID</key>
        <string>C218A47D-DAFB-4476-9C67-597E556D7D8A</string>
        <key>BucketName</key>
        <string>rsc</string>
        <key>ComputerUUID</key>
        <string>E7859547-BB9C-41C0-871E-858A0526BAE7</string>
        <key>LocalPath</key>
        <string>/Users/rsc</string>
        <key>LocalMountPoint</key>
        <string>/Users</string>
        <key>IgnoredRelativePaths</key>
        <array>
            <string>/.Trash</string>
            <string>/go/pkg</string>
            <string>/go1/pkg</string>
            <string>/Library/Caches</string>
        </array>
        <key>Dead</key>
        <string>foo<br/>baz</string>
        <key>Enabled</key>
        <true/>
        <key>Disabled</key>
        <false/>
        <key>IgnoredBool</key>
        <true/>
        <key>Excludes</key>
        <dict>
            <key>excludes</key>
            <array>
                <dict>
                    <key>type</key>
                    <integer>2</integer>
                    <key>text</key>
                    <string>.unison.</string>
                </dict>
            </array>
        </dict>
    </dict>
</plist>
`)

var xmlPrefix = []byte(`<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
`)

type TestStruct struct {
	BucketUUID           string
	BucketName           string
	ComputerUUID         string
	LocalPath            string
	LocalMountPoint      string
	IgnoredRelativePaths []string
	Excludes             Exclude
	Enabled              bool
	Disabled             bool
}

type Exclude struct {
	Excludes []ExcludeKey `plist:"excludes"`
}

type ExcludeKey struct {
	Type int    `plist:"type"`
	Text string `plist:"text"`
}

func TestUnmarshal(t *testing.T) {
	type args struct {
		data []byte
		v    interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "plist only",
			args: args{
				data: testPlist,
				v: &TestStruct{
					BucketUUID:      "C218A47D-DAFB-4476-9C67-597E556D7D8A",
					BucketName:      "rsc",
					ComputerUUID:    "E7859547-BB9C-41C0-871E-858A0526BAE7",
					LocalPath:       "/Users/rsc",
					LocalMountPoint: "/Users",
					IgnoredRelativePaths: []string{
						"/.Trash",
						"/go/pkg",
						"/go1/pkg",
						"/Library/Caches",
					},
					Enabled: true,
					Excludes: Exclude{
						Excludes: []ExcludeKey{
							{Type: 2,
								Text: ".unison.",
							},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "plist only",
			args: args{
				data: append(xmlPrefix, testPlist...),
				v: &TestStruct{
					BucketUUID:      "C218A47D-DAFB-4476-9C67-597E556D7D8A",
					BucketName:      "rsc",
					ComputerUUID:    "E7859547-BB9C-41C0-871E-858A0526BAE7",
					LocalPath:       "/Users/rsc",
					LocalMountPoint: "/Users",
					IgnoredRelativePaths: []string{
						"/.Trash",
						"/go/pkg",
						"/go1/pkg",
						"/Library/Caches",
					},
					Enabled: true,
					Excludes: Exclude{
						Excludes: []ExcludeKey{
							{Type: 2,
								Text: ".unison.",
							},
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := Unmarshal(tt.args.data, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal(%v, %v) error = %v, wantErr %v", tt.args.data, tt.args.v, err, tt.wantErr)
			}
			t.Logf("tt.args.v: %T => %+v\n", tt.args.v, tt.args.v)
		})
	}
}
