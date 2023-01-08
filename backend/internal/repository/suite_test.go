package repository

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestSuite(t *testing.T) {
	suite.Run(t, new(CardRepositoryTestSuite))
	suite.Run(t, new(DeckRepositoryTestSuite))
	suite.Run(t, new(MatchRepositoryTestSuite))
	suite.Run(t, new(RoundRepositoryTestSuite))
}
