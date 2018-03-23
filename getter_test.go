package main

import (
	"reflect"
	"testing"
)

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

		{
			source: []byte(`com.101tec
com.amazonaws
com.android.tools.lint
com.fasterxml`),
			filter: "all",
			expected: []string{
				"com.101tec",
				"com.amazonaws",
				"com.android.tools.lint",
				"com.fasterxml",
			},
		},
	}

	for _, testObject := range testObjects {
		actual := filterDir(testObject.source, testObject.filter)
		if !reflect.DeepEqual(actual, testObject.expected) {
			t.Errorf("expected = %+v, actual = %+v\n", testObject.expected, actual)
		}
	}
}

func TestGetArtifactName(t *testing.T) {
	testObjects := []struct {
		source   string
		expected string
	}{
		{
			source:   "com.helloworld.common.gradle/3.01/java.jar",
			expected: "com.helloworld.common.gradle/3.01/java",
		},
		{
			source:   "com.helloworld.common.gradle/3.01/java.pom",
			expected: "com.helloworld.common.gradle/3.01/java",
		},
		{
			source:   "com.helloworld.common.gradle/3.01/accessor-1.0.2.jar",
			expected: "com.helloworld.common.gradle/3.01/accessor-1.0.2",
		},
		{
			source:   "com.helloworld.common.gradle/3.01/accessor-1.0.2.pom",
			expected: "com.helloworld.common.gradle/3.01/accessor-1.0.2",
		},
	}

	for _, testObject := range testObjects {
		actual := getArtifactName(testObject.source)
		if actual != testObject.expected {
			t.Errorf("expected = %+v, actual = %+v\n", testObject.expected, actual)
		}
	}
}

func TestGetArtifactNameForGradle(t *testing.T) {
	testObjects := []struct {
		source   string
		expected string
	}{
		{
			source:   "com.helloworld.common.gradle/3.01/11111/java.jar",
			expected: "com.helloworld.common.gradle/3.01/java",
		},
		{
			source:   "com.helloworld.common.gradle/3.01/1231/java.pom",
			expected: "com.helloworld.common.gradle/3.01/java",
		},
		{
			source:   "com.helloworld.common.gradle/3.01/123123123/accessor-1.0.2.jar",
			expected: "com.helloworld.common.gradle/3.01/accessor-1.0.2",
		},
		{
			source:   "com.helloworld.common.gradle/3.01/123/accessor-1.0.2.pom",
			expected: "com.helloworld.common.gradle/3.01/accessor-1.0.2",
		},
	}

	for _, testObject := range testObjects {
		actual := getArtifactNameForGradle(testObject.source)
		if actual != testObject.expected {
			t.Errorf("expected = %+v, actual = %+v\n", testObject.expected, actual)
		}
	}
}
