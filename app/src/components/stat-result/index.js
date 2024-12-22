import React from 'react'
import PropTypes from 'prop-types'
import styled from 'styled-components'

import { t } from '../../theme'
import {
  Flex,
  Icon,
} from '../../ui'
import { moves } from '../../helpers'

export function StatResult({
  computerMove,
  playerMove,
}) {
  function icon(m) {
    return m !== undefined ? `hand-${moves[m]}` : 'question'
  }

  return (
    <Flex
      p={[5, 0, 5, 0]}
    >
      {[playerMove, computerMove].map((m, i) => (
        <StyledIcon
          c="white.5"
          key={i}
          fs={8}
          icon={icon(m)}
        />
      ))}
    </Flex>
  )
}

StatResult.propTypes = {
  computerMove: PropTypes.number,
  playerMove: PropTypes.number,
}

const StyledIcon = styled(Icon)`
  margin: 0 ${t('space.5')};
`
