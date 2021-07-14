import React from 'react'
import PropTypes from 'prop-types'
import styled from 'styled-components'

import { t } from '../../theme'
import { Flex } from '../../ui'

export function AppContent({
  children,
  header,
  footer,
}) {
  return (
    <StyledFlex
      align="stretch"
      content="flex-start"
      direction="column"
      p={[0, 4, 0, 4]}
    >
      {header && (
        <Header
          direction="column"
        >
          {header}
        </Header>
      )}
      <Main>
        {children}
      </Main>
      {footer && (
        <footer>
          {footer}
        </footer>
      )}
    </StyledFlex>
  )
}

AppContent.propTypes = {
  children: PropTypes.node.isRequired,
  header: PropTypes.node,
  footer: PropTypes.node,
}

const StyledFlex = styled(Flex)`
  height: 100vh;
  margin: 0 auto;
  max-width: ${t('size.sm')};
`

const Header = styled(Flex).attrs({
  as: 'header',
})`
  border-bottom: 1px solid ${t('color.white.5')};
`

const Main = styled.main`
  flex: 1;
`
