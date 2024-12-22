import React from 'react'
import PropTypes from 'prop-types'
import styled from 'styled-components'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { library } from '@fortawesome/fontawesome-svg-core'
import {
  faHandPaper,
  faHandRock,
  faHandScissors,
  faQuestion,
  faTrash,
  faUserNinja,
} from '@fortawesome/free-solid-svg-icons'

import { t } from '../../theme'

library.add(
  faHandPaper,
  faHandRock,
  faHandScissors,
  faQuestion,
  faTrash,
  faUserNinja,
)

export function Icon({
  icon,
  ...props
}) {
  return (
    <Wrapper { ...props }>
      <FontAwesomeIcon { ...{ icon } } />
    </Wrapper>
  )
}

Icon.propTypes = {
  c: PropTypes.string,
  fs: PropTypes.number,
  icon: PropTypes.string.isRequired,
}

Icon.defaultProps = {
  c: 'black.5',
  fs: 5,
}

const Wrapper = styled.span`
  color: ${({ c }) => t(`color.${c}`)};
  display: flex;
  font-size: ${({ fs }) => t(`fontSize.${fs}`)};
  justify-content: center;

  & > svg {
    fill: currentcolor;
    height: 100%;
    width: 100%;
  }
`
