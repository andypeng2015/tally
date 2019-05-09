// Copyright (c) 2019 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package tally

import (
	"testing"
	"time"
)

func BenchmarkSimpleCounterInc(b *testing.B) {
	c := &testSimpleCounter{}
	for n := 0; n < b.N; n++ {
		c.Inc(1)
	}
}

func BenchmarkSimpleCounterExpiredInc(b *testing.B) {
	c := &testSimpleCounter{}
	for n := 0; n < b.N; n++ {
		c.IncWithExpiredCheck(1)
	}
}
func BenchmarkAlwaysCheckCounterInc(b *testing.B) {
	s := newRootScope(ScopeOptions{Reporter: newTestStatsReporter()}, 0)
	c := newTestAlwaysCheckCounter("", s)
	for n := 0; n < b.N; n++ {
		c.Inc(1)
	}
}

func BenchmarkCounterInc(b *testing.B) {
	scope := NewTestScope("", nil)
	c := scope.Counter("test")
	for n := 0; n < b.N; n++ {
		c.Inc(1)
	}
}

func BenchmarkReportCounterNoData(b *testing.B) {
	c := &counter{}
	for n := 0; n < b.N; n++ {
		c.report("foo", nil, NullStatsReporter, 0)
	}
}

func BenchmarkReportCounterWithData(b *testing.B) {
	c := &counter{}
	for n := 0; n < b.N; n++ {
		c.Inc(1)
		c.report("foo", nil, NullStatsReporter, 0)
	}
}

func BenchmarkGaugeSet(b *testing.B) {
	g := &gauge{}
	for n := 0; n < b.N; n++ {
		g.Update(42)
	}
}

func BenchmarkReportGaugeNoData(b *testing.B) {
	g := &gauge{}
	for n := 0; n < b.N; n++ {
		g.report("bar", nil, NullStatsReporter, 0)
	}
}

func BenchmarkReportGaugeWithData(b *testing.B) {
	g := &gauge{}
	for n := 0; n < b.N; n++ {
		g.Update(73)
		g.report("bar", nil, NullStatsReporter, 0)
	}
}

func BenchmarkTimerStopwatch(b *testing.B) {
	t := &timer{
		name:     "bencher",
		tags:     nil,
		reporter: NullStatsReporter,
	}
	for n := 0; n < b.N; n++ {
		t.Start().Stop() // start and stop
	}
}

func BenchmarkTimerReport(b *testing.B) {
	t := &timer{
		name:     "bencher",
		tags:     nil,
		reporter: NullStatsReporter,
	}
	for n := 0; n < b.N; n++ {
		start := time.Now()
		t.Record(time.Since(start))
	}
}
