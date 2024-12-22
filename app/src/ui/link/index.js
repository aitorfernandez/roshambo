import React from 'react'
import PropTypes from 'prop-types'
import styled, {
  css,
} from 'styled-components'
import { NavLink } from 'react-router-dom'

import { t } from '../../theme'

export function Link({
  ...props
}) {
  if (props.to) {
    return <StyledNavLink { ...props } />
  }
  return <Anchor { ...props } />
}

Link.propTypes = {
  href: PropTypes.string,
  to: PropTypes.string,
  c: PropTypes.array,
  underline: PropTypes.number,
}

Link.defaultProps = {
  c: ['gray.7', 'gray.9'],
}

const styles = css`
  color: ${({ c }) => t(`color.${c[0]}`)};
  text-decoration: underline;
  transition: all ${t('transition.ease')};

  &:hover {
    color: ${({ c }) => t(`color.${c[1]}`)};
  }
`

const Anchor = styled.a`${styles}`

const StyledNavLink = styled(({
  ...props
}) => <NavLink { ...props } />)`${styles}`
