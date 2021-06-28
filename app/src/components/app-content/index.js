import React from 'react'
import PropTypes from 'prop-types'

export function AppContent({
  children,
  header,
}) {
  return (
    <div>
      {header && (
        <header>
          {header}
        </header>
      )}
      <main>
        {children}
      </main>
    </div>
  )
}

AppContent.propTypes = {
  children: PropTypes.node.isRequired,
  header: PropTypes.node,
}
