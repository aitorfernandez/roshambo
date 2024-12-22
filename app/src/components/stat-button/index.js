import React from 'react'
import styled from 'styled-components'

import { t } from '../../theme'
import {
  ButtonBase,
  Icon,
} from '../../ui'

export function StatButton({
  onClick,
  ...props
}) {
  return (
    <StyledButtonBase { ...{ onClick, ...props } }>
      <Icon { ...props } />
    </StyledButtonBase>
  )
}

StatButton.propTypes = {
  ...ButtonBase.propTypes,
  ...Icon.propTypes,
}

StatButton.defaultProps = {
  ...ButtonBase.defaultProps,
  ...Icon.defaultProps,
}

const StyledButtonBase = styled(ButtonBase)`
  background: ${t('color.white.5')};
  border-radius: 50%;
  height: 84px;
  transition: all ${t('transition.ease')};
  width: 84px;

  &:hover {
    transform: scale(1.2);
    background: ${t('color.white.5')};
  }
`
