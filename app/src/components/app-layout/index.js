import React from 'react'
import PropTypes from 'prop-types'
import { ThemeProvider } from 'styled-components'

import { AppGlobalStyle } from '../../AppGlobalStyle'
import { theme } from '../../theme'

export function AppLayout({
  children,
}) {
  return (
    <ThemeProvider { ...{ theme } }>
      <AppGlobalStyle />
      {children}
    </ThemeProvider>
  )
}

AppLayout.propTypes = {
  children: PropTypes.node.isRequired,
}
