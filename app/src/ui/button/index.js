import PropTypes from 'prop-types'
import styled from 'styled-components'

import { ButtonBase } from '../button-base'
import { t } from '../../theme'

export const Button = styled(ButtonBase)`
  background-color: ${({ bg }) => t(`color.${bg[0]}`)};
  border-radius: ${t(`borderRadius.3`)};
  color: ${({ c }) => t(`color.${c}`)};
  font-size: ${({ fs}) => t(`fontSize.${fs}`)};
  opacity: ${({ disabled }) => t(`opacity.${disabled ? 2 : 5}`)};
  padding: ${t('space.5')} ${t('space.7')};
  transition: ${t('transition.ease')};

  &:hover {
    background-color: ${({ disabled, bg }) => t(`color.${disabled ? bg[0] : bg[1]}`)};
  }
`

Button.propTypes = {
  ...ButtonBase.propTypes,
  bg: PropTypes.array,
  c: PropTypes.string,
  disabled: PropTypes.bool,
  fs: PropTypes.number,
}

Button.defaultProps = {
  bg: ['gray.5', 'gray.7'],
  c: 'white.5',
  disabled: false,
  fs: 1,
}
