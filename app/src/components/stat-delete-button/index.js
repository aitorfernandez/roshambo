import React from 'react'
import PropTypes from 'prop-types'
import { useCookies } from 'react-cookie'

export function StatDeleteButton({
  deleteStats,
}) {
  const [{ id: accountID }] = useCookies(['id'])

  async function onSubmit() {
    await deleteStats({
      variables: {
        accountID,
      },
    })
  }

  return (
    <div>
      <button
        onClick={onSubmit}
      >
        StatDeleteButton
      </button>
    </div>
  )
}

StatDeleteButton.propTypes = {
  deleteStats: PropTypes.func.isRequired,
}
