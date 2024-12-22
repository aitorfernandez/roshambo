import { PlayPageAccountQuery } from '../containers/play-page'
import {
  readQuery,
  writeQuery,
} from './helpers'

function addStat(cache, stat, vars) {
  const query = readQuery(cache, PlayPageAccountQuery, vars)
  if (!query) {
    return
  }
  const data = {
    account: {
      ...query.account,
      stats: [stat, ...query.account.stats],
    },
  }
  writeQuery(cache, PlayPageAccountQuery, data, vars)
}

function addProfile(cache, profile, vars) {
  const query = readQuery(cache, PlayPageAccountQuery, vars)
  if (!query) {
    return
  }
  const data = {
    account: {
      ...query.account,
      profile,
    },
  }
  writeQuery(cache, PlayPageAccountQuery, data, vars)
}

export const fromAccount = {
  addProfile,
  addStat,
}
