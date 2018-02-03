package main

import (
	"reflect"
	"testing"
)

func TestGetFolder(t *testing.T) {
	testObjects := []struct {
		source   string
		expected []string
	}{
		{
			source:   "com.helloworld.common.gradle",
			expected: []string{"com", "helloworld", "common", "gradle"},
		},

		{
			source:   "com.helloworld.common.gradle/library",
			expected: []string{"com", "helloworld", "common", "gradle", "library"},
		},

		{
			source:   "com.hello-world.common.gradle/library/java-plugin/4.10.0",
			expected: []string{"com", "hello-world", "common", "gradle", "library", "java-plugin", "4.10.0"},
		},
	}

	for _, testObject := range testObjects {
		actual := getFolder(testObject.source)
		if !reflect.DeepEqual(actual, testObject.expected) {
			t.Errorf("expected = %+v, actual = %+v\n", testObject.expected, actual)
		}
	}
}

func TestFilterDir(t *testing.T) {
	testObjects := []struct {
		source   []byte
		filter   string
		expected []string
	}{
		{
			source: []byte(`com.101tec
com.amazonaws
com.amazon.redshift
com.android.databinding
com.android.tools.jill
com.android.tools.layoutlib
com.android.tools.lint
com.auth0
com.beust
com.boundary
com.clearspring.analytics
com.cloudbees
com.cloudbees.thirdparty
com.codahale.metrics
com.codeborne
com.cybozu.labs
com.damnhandy
com.datadoghq
com.datastax.cassandra
com.ecyrd.speed4j
com.esotericsoftware.kryo
com.esotericsoftware.minlog
com.esotericsoftware.reflectasm
com.factual
com.fasterxml`),
			filter:   "android",
			expected: []string{"com.android.databinding", "com.android.tools.jill", "com.android.tools.layoutlib", "com.android.tools.lint"},
		},
	}

	for _, testObject := range testObjects {
		actual := filterDir(testObject.source, testObject.filter)
		if !reflect.DeepEqual(actual, testObject.expected) {
			t.Errorf("expected = %+v, actual = %+v\n", testObject.expected, actual)
		}
	}
}
