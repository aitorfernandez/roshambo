import PropTypes from 'prop-types'
import styled from 'styled-components'

export const ButtonBase = styled.button.attrs(({
  disabled,
  onClick,
}) => ({
  onClick: disabled ? undefined : onClick,
}))`
  background-color: inherit;
  border: none;
  color: inherit;
  cursor: pointer;
  display: inline-block;
  font: inherit;
  margin: 0;
  overflow: hidden;
  text-align: center;
  text-decoration: none;
  vertical-align: middle;
  white-space: nowrap

  &:hover {
    text-decoration: none;
  }

  &:focus {
    outline: none;
  }

  &:disabled {
    cursor: not-allowed;
  }
`

ButtonBase.propTypes = {
  onClick: PropTypes.func,
  type: PropTypes.string,
}

ButtonBase.defaultProps = {
  type: 'button',
}
