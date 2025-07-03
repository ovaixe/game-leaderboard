FROM node:20-alpine

ARG UID=1000
ARG GID=1000

# Install system dependencies if needed
# RUN apk add --no-cache <your-dependencies>

# Create app directory
WORKDIR /app

# # Create a non-root user with the specified UID/GID
# RUN addgroup -g $GID appuser && \
#     adduser -u $UID -G appuser -D -H -s /bin/sh appuser

# Change ownership of the WORKDIR to appuser
# RUN chown -R appuser:appuser /app

# Switch to appuser
# USER appuser

# Copy package files first to leverage Docker cache
COPY package.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application code
COPY . .

# Build the application
RUN npm run build

# Health check (optional)
# HEALTHCHECK --interval=30s --timeout=3s \
#   CMD curl -f http://localhost:5173/ || exit 1

EXPOSE 5173

CMD ["npm", "run", "dev", "--", "--host", "0.0.0.0"]