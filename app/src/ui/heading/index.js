import React from 'react'
import PropTypes from 'prop-types'
import styled from 'styled-components'

import { t } from '../../theme'

const fs = ({
  level,
  ...props
}) => t(`fontSize.${7 - level}`)(props)

export const Heading = styled(({
  children,
  level,
  ...props
}) => React.createElement(`h${level}`, props, children))`
  font-size: ${fs};
  font-weight: ${t('fontWeight.bold')};
  letter-spacing: normal;
  line-height: 1.2;
`

Heading.propTypes = {
  level: PropTypes.number.isRequired,
}

Heading.defaultProps = {
  level: 1,
}
