{
  description = "Ambiente de ferramentas de Infraestrutura";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
  };

  outputs = { self, nixpkgs }:
    let
      system = "x86_64-linux";
      pkgs = nixpkgs.legacyPackages.${system};
    in
    {
      devShells.${system}.default = pkgs.mkShell {
        buildInputs = with pkgs; [
          go
          gopls
          delve # for debug go

          nodejs
          
          #docker-compose
          rabbitmq-c # Bibliotecas C para rabbitmq se precisar
          # se quiser ferramentas extras como elasticdump, pode por aqui
        ];
        
        hardeningDisable = [ "fortify" ];

        shellHook = ''
          mkdir -p .node_modules_global/bin
          mkdir -p .node_modules_global/cache

          export NPM_CONFIG_PREFIX=$PWD/.node_modules_global
          export PATH=$PWD/.node_modules_global/bin:$PATH
          
          export NPM_CONFIG_CACHE=$PWD/.node_modules_global/cache

          mkdir -p .go/path/bin
          mkdir -p .go/cache

          export GOPATH=$PWD/.go/path
          export GOCACHE=$PWD/.go/cache

          export PATH=$GOPATH/bin:$PATH

          echo "✦ Ambiente de Infraestrutura Pronto ✦"
  
          testar-prod() {
            echo "🚀 1. Derrubando ambientes antigos..."
            podman-compose -f docker-compose.prod.yml down -v
            
            podman rm -f $(podman ps -a -q --filter name=infra_) 2>/dev/null || true
            
            echo "📦 2. Buildando os containers de produção com as alterações recentes..."
            podman-compose -f docker-compose.prod.yml build

            echo "🌐 3. Subindo toda a infraestrutura (SaaS + Banco + Mensageria + Logs)..."
            podman-compose -f docker-compose.prod.yml up -d
            
            echo "⏳ 4. Aguardando os serviços iniciarem completamente..."
            sleep 10 # Dá um tempo para o Postgres e o Elastic subirem
            
            echo "🔍 5. Executando testes automatizados de integração..."
            # Aqui você poderia rodar uma suite de testes do Playwright/Cypress
            # ou simplesmente fazer um curl de saúde:
            curl -f http://localhost:8080/health || echo "❌ Backend falhou no teste de saúde!"
            
            echo "📊 Verifique os logs no Kibana em http://localhost:5601"
            echo "🍔 Acesse a aplicação simulada em http://localhost:3000"
          }

          alias stop-prod='podman-compose -f docker-compose.prod.yml down -v'

          testar-dev() {
            echo "🚀 1. Derrubando ambientes antigos..."
            podman-compose -f docker-compose.yml down -v
            
            podman rm -f $(podman ps -a -q --filter name=infra_) 2>/dev/null || true
            
            echo "📦 2. Buildando os containers de produção com as alterações recentes..."
            podman-compose -f docker-compose.yml build

            echo "🌐 3. Subindo toda a infraestrutura (SaaS + Banco + Mensageria + Logs)..."
            podman-compose -f docker-compose.yml up -d
            
            echo "⏳ 4. Aguardando os serviços iniciarem completamente..."
            sleep 10 # Dá um tempo para o Postgres e o Elastic subirem
            
            echo "🔍 5. Executando testes automatizados de integração..."
            # Aqui você poderia rodar uma suite de testes do Playwright/Cypress
            # ou simplesmente fazer um curl de saúde:
            curl -f http://localhost:8080/health || echo "❌ Backend falhou no teste de saúde!"
            
            echo "📊 Verifique os logs no Kibana em http://localhost:5601"
            echo "🍔 Acesse a aplicação simulada em http://localhost:3000"
          }

          alias stop-dev='podman-compose -f docker-compose.yml down -v'
        '';
      };
    };
}