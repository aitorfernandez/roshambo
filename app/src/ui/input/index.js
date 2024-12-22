import React, {
  forwardRef,
} from 'react'
import PropTypes from 'prop-types'
import styled from 'styled-components'

import { t } from '../../theme'

export const Input = forwardRef(({ ...props}, ref) => {
  const { type } = props
  switch (type) {
    case 'textarea':
    case 'select':
      console.error('not implemented')
      break
    default:
      return <StyledInput ref={ref} { ...props }/>
  }
})

Input.defaultProps = {
  type: 'text',
}

Input.propTypes = {
  type: PropTypes.string.isRequired,
}

const StyledInput = styled.input`
  border-radius: ${t('borderRadius.3')};
  border: none;
  box-shadow: none;
  box-sizing: border-box;
  font-size: ${({ fs }) => t(`fontSize.${fs}`)};
  letter-spacing: 0.125px;
  line-height: 14px;
  padding: ${t('space.5')};
  width: 100%;

  &:active,
  &:focus {
    outline: currentcolor none medium;
  }
`

StyledInput.propTypes = {
  fs: PropTypes.number,
}

StyledInput.defaultProps = {
  fs: 2,
}
