import React from 'react'
import styled from 'styled-components'

import { t } from '../../theme'
import {
  Flex,
  Icon,
  Link,
} from '../../ui'

export function Header() {
  return (
    <Link to="/">
      <Flex>
        {['rock', 'paper', 'scissors'].map((item) => (
          <StyledIcon
            c="white.5"
            fs={2}
            icon={`hand-${item}`}
            key={item}
          />
        ))}
      </Flex>
    </Link>
  )
}

const StyledIcon = styled(Icon)`
  margin: 0 ${t('space.2')};
`
