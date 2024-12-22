import React from 'react'
import PropTypes from 'prop-types'
import styled from 'styled-components'
import { useCookies } from 'react-cookie'

import { Flex } from '../../ui'
import { moves } from '../../helpers'
import { StatButton } from '../stat-button'

export function StatCreateButton({
  createStat,
  playerMove,
}) {
  const [{ id: accountID }] = useCookies(['id'])

  async function onSubmit(values) {
    await createStat({
      variables: {
        input: {
          accountID,
          ...values,
        },
      },
    })
  }

  return (
    <StyledFlex
      p={[6, 0, 6, 0]}
    >
      <StatButton
        c="gray.7"
        fs={6}
        icon={`hand-${moves[playerMove]}`}
        onClick={() => onSubmit({ playerMove })}
      />
    </StyledFlex>
  )
}

StatCreateButton.propTypes = {
  createStat: PropTypes.func.isRequired,
  playerMove: PropTypes.number.isRequired,
}

const StyledFlex = styled(Flex)`
  flex: 1 1 auto;
`
