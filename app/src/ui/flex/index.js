import PropTypes from 'prop-types'
import styled from 'styled-components'

import { Box } from '../box'

export const Flex = styled(Box)`
  align-items: ${({ align }) => align};
  display: flex;
  flex-direction: ${({ direction }) => direction};
  justify-content: ${({ content }) => content};
`

Flex.propTypes = {
  ...Box.propTypes,
  align: PropTypes.string,
  content: PropTypes.string,
  direction: PropTypes.string,
}

Flex.defaultProps = {
  ...Box.defaultProps,
  align: 'center',
  content: 'center',
  direction: 'row',
}
