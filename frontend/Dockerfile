FROM node:20-alpine AS builder
WORKDIR /app

ARG VITE_API_URL_USERS
ARG VITE_API_URL_TRAINS
ARG VITE_API_URL_TICKETS
ARG VITE_API_URL_NOTIFICATIONS
ARG VITE_API_URL_AUTH

ENV VITE_API_URL_USERS=$VITE_API_URL_USERS
ENV VITE_API_URL_TRAINS=$VITE_API_URL_TRAINS
ENV VITE_API_URL_TICKETS=$VITE_API_URL_TICKETS
ENV VITE_API_URL_NOTIFICATIONS=$VITE_API_URL_NOTIFICATIONS
ENV VITE_API_URL_AUTH=$VITE_API_URL_AUTH

COPY package*.json ./
RUN npm ci --ignore-scripts
COPY . .
RUN npm run build

FROM nginx:alpine
COPY --from=builder /app/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY print-url.sh /print-url.sh

RUN sed -i 's/\r$//' /print-url.sh || true && \
    chmod +x /print-url.sh

RUN chown -R nginx:nginx /usr/share/nginx/html && \
    chown -R nginx:nginx /var/cache/nginx && \
    chown -R nginx:nginx /var/log/nginx && \
    chown -R nginx:nginx /etc/nginx/conf.d && \
    touch /var/run/nginx.pid && \
    chown nginx:nginx /var/run/nginx.pid && \
    chmod +x /print-url.sh

USER nginx

EXPOSE 80
CMD ["/print-url.sh"]