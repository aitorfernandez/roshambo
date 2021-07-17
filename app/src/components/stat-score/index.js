import React from 'react'
import PropTypes from 'prop-types'
import styled from 'styled-components'

import {
  Flex,
  Paragraph,
} from '../../ui'
import { t } from '../../theme'

export function StatScore({
  stats,
}) {
  let [win, lose, draw] = Array(3).fill(0)

  function winner(p, c) {
    if (p === c) {
      draw++
      return
    }
    if ((p + 1) % 3 === c) {
      lose++
      return
    }
    win ++
  }

  stats.forEach(({ computerMove, playerMove}) => winner(playerMove, computerMove))

  return (
    <Flex
      p={[1, 0, 1, 0]}
    >
      <StyledParagraph>
        <span>
          win <Bold>{win}</Bold>
        </span>
        <span>
          lose <Bold>{lose}</Bold>
        </span>
        <span>
          draw <Bold>{draw}</Bold>
        </span>
      </StyledParagraph>
    </Flex>
  )
}

StatScore.propTypes = {
  stats: PropTypes.arrayOf(
    PropTypes.shape({
      computerMove: PropTypes.number.isRequired,
      playerMove: PropTypes.number.isRequired,
    }).isRequired,
  ),
}

const StyledParagraph = styled(Paragraph)`
  margin: 0;

  span:nth-child(2) {
    margin: 0 ${t('space.7')};
  }
`

const Bold = styled.b`
  font-size: ${t('fontSize.4')};
  font-weight: ${t('fontWeight.bold')};
`
