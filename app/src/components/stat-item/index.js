import React from 'react'
import PropTypes from 'prop-types'
import { gql } from '@apollo/client'

export function StatItem({
  computerMove,
  playerMove,
  round,
}) {
  return (
    <div>
      {`${round} ${playerMove} ${computerMove}`}
    </div>
  )
}

StatItem.propTypes = {
  computerMove: PropTypes.number.isRequired,
  playerMove: PropTypes.number.isRequired,
  round: PropTypes.number.isRequired,
}

StatItem.fragments = {
  stat: gql`
    fragment StatItemStatFragment on Stat {
      id
      computerMove
      playerMove
      round
    }
  `,
}
