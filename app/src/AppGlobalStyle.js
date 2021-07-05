import { createGlobalStyle } from 'styled-components'
import PropTypes from 'prop-types'

import { t } from './theme'

export const AppGlobalStyle = createGlobalStyle`
  *,
  *::before,
  *::after {
    box-sizing: border-box;
  }

  html {
    -ms-text-size-adjust: 100%;
    -webkit-text-size-adjust: 100%;
    font-family: sans-serif;
  }

  html, body {
    height: 100%;
  }

  html {
    font-size: ${t('fontBaseSize')}px;
  }

  body {
    -moz-font-feature-settings: 'liga' on;
    -moz-osx-font-smoothing: grayscale;
    -webkit-font-smoothing: antialiased;
    line-height: 1.5;
    margin: 0;
    text-rendering: optimizeLegibility;
    word-wrap: break-word;
  }

  body {
    background: ${({ bg }) => t(`color.${bg}`)};
    color: ${({ c }) => t(`color.${c}`)};
    font-family: ${({ font }) => t(`fontFamily.${font}`)};
  }
`

AppGlobalStyle.propTypes = {
  bg: PropTypes.string,
  c: PropTypes.string,
  fontFamily: PropTypes.string,
}

AppGlobalStyle.defaultProps = {
  bg: 'white.5',
  c: 'black.5',
  font: 'primary',
}
