import PropTypes from 'prop-types'
import styled from 'styled-components'

import { t } from '../../theme'

export const Paragraph = styled.p`
  font-size: ${({ fs }) => t(`fontSize.${fs}`)};
  letter-spacing: 0.024rem;
  line-height: 1.4;
  text-align: ${({ ta }) => ta};
`

Paragraph.propTypes = {
  fs: PropTypes.number,
  ta: PropTypes.string,
}

Paragraph.defaultProps = {
  fs: 1,
  ta: 'left',
}
