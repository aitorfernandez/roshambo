import PropTypes from 'prop-types'
import styled from 'styled-components'

import { t } from '../../theme'

function space(s, props) {
  return s.map((e) => t(`space.${e}`)(props)).join(' ')
}

export const Box = styled.div`
  font-size: ${({ fs }) => t(`fontSize.${fs}`)};
  margin: ${({ m, ...props }) => space(m, props)};
  padding: ${({ p, ...props }) => space(p, props)};
`

Box.propTypes = {
  fs: PropTypes.number,
  m: PropTypes.array,
  p: PropTypes.array,
}

Box.defaultProps = {
  fs: 1,
  m: [0, 0, 0, 0],
  p: [0, 0, 0, 0],
}
