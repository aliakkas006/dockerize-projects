# === Build Stage ===
FROM node:20-alpine AS builder

WORKDIR /app

# Copy dependencies first for caching
COPY package.json yarn.lock ./

# Install only dependencies
RUN yarn install --frozen-lockfile

# Copy source code
COPY . .

# Build Next.js project
RUN yarn build

# === Production Stage ===
FROM node:20-alpine AS runner

WORKDIR /app

# Copy only necessary files for production runtime
COPY package.json yarn.lock ./
COPY --from=builder /app/.next .next
COPY --from=builder /app/public ./public
COPY --from=builder /app/node_modules ./node_modules
COPY --from=builder /app/next.config.js ./next.config.js

EXPOSE 3000

CMD ["yarn", "start"]
