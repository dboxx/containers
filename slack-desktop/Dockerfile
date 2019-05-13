FROM debian:buster-slim

# OS
ENV DEBIAN_FRONTEND "noninteractive"

# Slack configuration
ENV SLACK_VERSION      "3.4.0"
ENV SLACK_ARCH         "amd64"

# Slack download configuration
ENV SLACK_DOWNLOAD_PKG "slack-desktop-${SLACK_VERSION}-${SLACK_ARCH}.deb"
ENV SLACK_DOWNLOAD_URL "https://downloads.slack-edge.com/linux_releases"
ENV SLACK_DOWNLOAD_DIR "/tmp"

ADD "${SLACK_DOWNLOAD_URL}/${SLACK_DOWNLOAD_PKG}" "${SLACK_DOWNLOAD_DIR}"

RUN apt-get update -y \
 && apt-get install -yy -q libxss1 \
                           libappindicator1 \
                           libcurl4 \
                           libsecret-1-0 \
                           gconf2 \
                           gconf-service \
                           libgtk2.0-0 \
                           libnotify4 \
                           libxtst6 \
                           libnss3 \
                           python \
                           gvfs-bin \
                           xdg-utils \
                           libasound2 \
                           libasound2-plugins \
 && dpkg -i "${SLACK_DOWNLOAD_DIR}/${SLACK_DOWNLOAD_PKG}" \
 && apt-get autoremove -y --purge \
 && apt-get autoclean -y \
 && apt-get clean -y \
 && apt-get check -y \
 && rm -rf "${SLACK_DOWNLOAD_DIR}/*" /var/lib/apt/lists/* /var/cache/* /usr/src/* /tmp/* /var/tmp/*

ENV DEBIAN_FRONTEND "noninteractive"

ENTRYPOINT slack
