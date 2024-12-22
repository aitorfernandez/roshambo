import React from 'react'
import PropTypes from 'prop-types'
import styled from 'styled-components'
import { useCookies } from 'react-cookie'

import { Flex } from '../../ui'
import { StatButton } from '../stat-button'

export function StatDeleteButton({
  deleteStats,
}) {
  const [{ id: accountID }] = useCookies(['id'])

  async function onSubmit() {
    await deleteStats({
      variables: {
        accountID,
      },
    })
  }

  return (
    <StyledFlex>
      <StatButton
        c="red.5"
        fs={6}
        icon="trash"
        onClick={onSubmit}
      />
    </StyledFlex>
  )
}

StatDeleteButton.propTypes = {
  deleteStats: PropTypes.func.isRequired,
}

const StyledFlex = styled(Flex)`
  flex: 1 1 auto;
`
