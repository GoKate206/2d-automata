package main

import (
	"testing"
  "reflect"

	"github.com/stretchr/testify/assert"
)

func TestPopulateGeneration(t *testing.T) {
  gridA := Grid{
    {0,0,2,2,0,0,0,0,0,0},
    {0,0,0,0,1,3,1,0,0,0},
    {0,0,0,2,0,3,0,0,1,3},
    {0,0,0,0,1,3,0,0,0,3},
    {0,2,2,0,0,0,1,3,1,0},
    {0,0,0,0,0,2,2,0,0,0},
    {0,0,2,0,2,0,0,0,0,0},
    {0,0,0,0,2,0,0,0,0,0},
    {0,0,0,0,0,0,0,0,0,0},
    {0,0,0,0,0,0,0,0,0,0},
  }

  gridB := Grid{
    {0,0,0,0,0,0,0,0,0,0},
    {0,0,1,1,0,0,2,0,0,0},
    {0,0,3,3,2,0,0,0,2,0},
    {0,1,0,0,0,0,0,0,2,0},
    {0,0,3,0,0,1,2,0,0,0},
    {0,0,1,3,3,3,0,0,0,0},
    {0,0,0,1,0,1,0,0,0,0},
    {0,0,0,0,0,0,0,0,0,0},
    {0,0,0,0,0,0,0,0,0,0},
    {0,0,0,0,0,0,0,0,0,0},
  }


  t.Run("when populateNewGeneration is invoked with valid gridA", func(t *testing.T) {
    result := populateNewGeneration(gridA)
    t.Run("it should match expected", func(t *testing.T) {
      expected := Grid{
        {0,0,3,3,0,0,0,0,0,0},
        {0,0,0,0,2,0,2,0,0,0},
        {0,0,0,3,0,0,0,0,2,0},
        {0,1,0,1,2,0,0,0,0,0},
        {0,3,3,0,0,1,2,0,2,0},
        {0,0,0,0,1,0,0,0,0,0},
        {0,0,0,0,3,0,1,0,0,0},
        {0,0,0,0,3,1,0,0,0,0},
        {0,0,0,0,0,0,0,0,0,0},
        {0,0,0,0,0,0,0,0,0,0},
      }
      isEqual := reflect.DeepEqual(result, expected)
      assert.True(t, isEqual)
    })
  })

  t.Run("when populateNewGeneration is invoked with valid gridB", func(t *testing.T) {
    result := populateNewGeneration(gridB)
    t.Run("it should match expected", func(t *testing.T) {
      expected := Grid{
        {0,0,0,0,0,0,0,0,0,0},
        {0,0,2,2,0,1,0,1,0,0},
        {0,0,0,0,3,1,0,0,3,1},
        {0,2,0,0,0,1,0,0,3,1},
        {0,0,0,0,0,2,3,1,0,0},
        {0,0,2,0,0,0,0,0,0,0},
        {0,0,0,2,0,2,0,0,0,0},
        {0,0,0,0,0,0,0,0,0,0},
        {0,0,0,0,0,0,0,0,0,0},
        {0,0,0,0,0,0,0,0,0,0},
      }
      isEqual := reflect.DeepEqual(result, expected)
      assert.True(t, isEqual)
    })
  })
}
