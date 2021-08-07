#!/usr/bin/env bash

# Shell script to build my blog

set -euo pipefail;

pages=${CF_PAGES:-0}

# If not running in Cloudflare workers, just run `zola serve`

if [[ "${pages}" == 0 ]]; then
  zola serve;
  exit 0
fi

if [[ -z "${CF_PAGES_BRANCH}" ]]; then

  if [[ "${CF_PAGES_BRANCH}" == "main" ]]; then
    zola build
  else

    if [[ -z "${CF_PAGES_COMMIT_SHA}" ]]; then
      zola build --base-url "https://${CF_PAGES_COMMIT_SHA}.sp-zola-blog.pages.dev"
    else
      echo "env CF_PAGES_COMMIT_SHA not set!";
      exit 1
    fi
  fi

else
  echo "env CF_PAGES_BRANCH not set!";
  exit 1
fi

