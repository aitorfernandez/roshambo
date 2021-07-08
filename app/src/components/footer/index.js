import React from 'react'
import styled from 'styled-components'

import { t } from '../../theme'
import { Box } from '../../ui'

export function Footer() {
  return (
    <StyledBox
      fs={0}
      m={[8, 0, 0, 0]}
      p={[5, 0, 5, 0]}
    >
      Thanks for playing ＼(＾▽＾)／
    </StyledBox>
  )
}

const StyledBox = styled(Box)`
  border-top: 1px solid ${t('color.white.5')};
  text-align: center;
`
