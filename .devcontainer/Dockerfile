FROM denoland/deno:bin-2.4.1 AS deno
FROM mcr.microsoft.com/devcontainers/universal
COPY --from=deno /deno /usr/local/bin/deno
