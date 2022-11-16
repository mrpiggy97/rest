# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.234.0/containers/go/.devcontainer/base.Dockerfile

# [Choice] Go version (use -bullseye variants on local arm64/Apple Silicon): 1, 1.16, 1.17, 1-bullseye, 1.16-bullseye, 1.17-bullseye, 1-buster, 1.16-buster, 1.17-buster
ARG VARIANT="1.18-bullseye"
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}

# [Choice] Node.js version: none, lts/*, 16, 14, 12, 10
ARG NODE_VERSION="none"
ENV PORT=5000
ENV TESTING_DATABASE_URL=postgres://testing_user:testing_password@testingdb:5432/restdb_testing?sslmode=disable
ENV TESTING_JWT_SECRET=T4zcWNdCN8MsWcTVI7u6noSg8a9F8SbS9Pu3pnnv/GcXDi5UwC/Pdnm3eYeoeVQEGWBXVjlWgaMHroTIoO6Rwk4LENknD8nL3vtlT/ewjRQC4AK8SJ0NVBXjQOZ0SyG0n3Cgd79/d3QABC/UQhssgRD8652wHPvKBWKufIwy6pqvIGK9eCpWv0SFkWXb7NmA0408XCzVpcspqqOnjCBzOLIJrHrEEY70JYqcOP0DPEkkEf0pkCBBgtY9E9lJgnwGGtDi6YVn3dYD0PpKiWwf2gMCu0aj1+7/B9ESY9zIr3Jum/cSQs2sq3kwUke8w9kNotiDdnNvj0GsRJ087nyVPQ==
RUN if [ "${NODE_VERSION}" != "none" ]; then su vscode -c "umask 0002 && . /usr/local/share/nvm/nvm.sh && nvm install ${NODE_VERSION} 2>&1"; fi

# [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends <your-package-list-here>

# [Optional] Uncomment the next lines to use go get to install anything else you need
# USER vscode
# RUN go get -x <your-dependency-or-tool>

# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1
