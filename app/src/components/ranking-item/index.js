import React from 'react'
import PropTypes from 'prop-types'
import { gql } from '@apollo/client'

export function RankingItem({
  draw,
  lose,
  totalRounds,
  username,
  win,
}) {
  return (
    <div>
      {`${username} ${totalRounds} ${win} ${lose} ${draw}`}
    </div>
  )
}

RankingItem.propTypes = {
  draw: PropTypes.number.isRequired,
  lose: PropTypes.number.isRequired,
  totalRounds: PropTypes.number.isRequired,
  username: PropTypes.string.isRequired,
  win: PropTypes.number.isRequired,
}

RankingItem.fragments = {
  ranking: gql`
    fragment RankingItemRankingFragment on Ranking {
      id
      draw
      lose
      totalRounds
      username
      win
    }
  `,
}
