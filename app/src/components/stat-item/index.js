import React from 'react'
import PropTypes from 'prop-types'
import styled from 'styled-components'
import { gql } from '@apollo/client'

import { moves } from '../../helpers'
import { t }  from '../../theme'
import {
  Flex,
  Box,
  Icon,
} from '../../ui'

export function StatItem({
  computerMove,
  playerMove,
  round,
}) {
  function winner(p, c) {
    if (p === c) {
      return 'yellow'
    }
    if ((p + 1) % 3 === c) {
      return 'red'
    }
    return 'green'
  }

  function color(prefColor) {
    const c = winner(playerMove, computerMove)
    if (c !== 'yellow' && c !== prefColor) {
      return 'gray.5'
    }
    return `${c}.5`
  }

  return (
    <StyledFlex
      m={[4, 0, 4, 0]}
      p={[4, 0, 4, 0]}
    >
      <StyledBox
        fontSize={1}
      >
        ({round})
      </StyledBox>
      <StyledIcon
        c={color('green')}
        fs={4}
        icon={`hand-${moves[playerMove]}`}
      />
      <StyledIcon
        c={color('red')}
        fs={4}
        icon={`hand-${moves[computerMove]}`}
      />
    </StyledFlex>
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

const StyledFlex = styled(Flex)`
  background: ${t('color.white.5')};
  border-radius: ${t('borderRadius.2')};
`

const StyledBox = styled(Box)`
  color: ${t('color.gray.5')};
  letter-spacing: 0.024rem;
`

const StyledIcon = styled(Icon)`
  margin: 0 ${t('space.3')};
`
