// Copyright 2015 ThoughtWorks, Inc.

// This file is part of Gauge.

// Gauge is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// Gauge is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with Gauge.  If not, see <http://www.gnu.org/licenses/>.

package run_failed

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/getgauge/common"
	"github.com/getgauge/gauge/config"
	"github.com/getgauge/gauge/execution/result"
	"github.com/getgauge/gauge/gauge_messages"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestIfFailedFileIsCreated(c *C) {
	p, _ := filepath.Abs("_testdata")
	config.ProjectRoot = p
	failedInfo := "hello world"

	writeFailedMeta(failedInfo)

	file := filepath.Join(config.ProjectRoot, dotGauge, failedFile)
	c.Assert(common.FileExists(file), Equals, true)
	expected := failedInfo

	content, _ := ioutil.ReadFile(file)

	c.Assert(string(content), Equals, expected)
	os.RemoveAll(filepath.Join(config.ProjectRoot, dotGauge))
}

func (s *MySuite) TestGetFailedMetadata(c *C) {
	p, _ := filepath.Abs("_testdata")
	config.ProjectRoot = p
	spec1Rel := filepath.Join("specs", "example1.spec")
	spec1Abs := filepath.Join(p, spec1Rel)
	sr1 := &result.SpecResult{IsFailed: true, ProtoSpec: &gauge_messages.ProtoSpec{FileName: &spec1Abs}, FailedScenarioIndices: []int{2, 6}}

	meta := getFailedMetadata([]*result.SpecResult{sr1})

	c.Assert(len(meta.FailedScenarios), Equals, 2)
	c.Assert(meta.FailedScenarios[0], Equals, spec1Rel+":2")
	c.Assert(meta.FailedScenarios[1], Equals, spec1Rel+":6")
}
