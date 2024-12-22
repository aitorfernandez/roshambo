import React, {
  useRef,
  useState,
} from 'react'
import PropTypes from 'prop-types'
import styled from 'styled-components'
import DOMPurify from 'dompurify'

import { t } from '../../theme'

export function ContentEditable({
  defaultValue,
  onEdit,
  value: v,
  ...props
}) {
  const [value, setValue] = useState(v || defaultValue)

  const boxRef = useRef(null)
  const valRef = useRef(value)

  function onBlur() {
    onEdit && onEdit(DOMPurify.sanitize(value))
  }

  function onClick(e) {
    if (e.target.innerText === defaultValue) {
      e.target.innerText = ''
    }
  }

  function onInput(e) {
    setValue(e.target.innerText)
  }

  function onKeyPress(e) {
    if (e.which === 13) {
      e.preventDefault()
      boxRef.current.blur()
    }
  }

  function onPaste(e) {
    e.preventDefault()
    return
  }

  return (
    <Wrapper
      contentEditable
      dangerouslySetInnerHTML={{ __html: valRef.current }}
      onBlur={onBlur}
      onClick={onClick}
      onInput={onInput}
      onKeyPress={onKeyPress}
      onPaste={onPaste}
      ref={boxRef}
      {...props}
    />
  )
}

ContentEditable.propTypes = {
  bg: PropTypes.string,
  defaultValue: PropTypes.string,
  onEdit: PropTypes.func,
  value: PropTypes.string,
}

ContentEditable.defaultProps = {
  bg: 'gray.5',
  defaultValue: 'content here...',
}

const Wrapper = styled.div`
  border-radius: ${t('borderRadius.2')};
  color: inherit;
  display: block;
  font-size: inherit;
  margin: 0;
  min-width: 0;
  outline: currentcolor none medium;
  overflow: hidden;
  padding: ${t('space.2')};
  position: relative;
  resize: none;
  word-break: break-word;

  &:focus,
  &:focus:hover,
  &:hover {
    background: ${({ bg }) => t(`color.${bg}`)};
  }
`
