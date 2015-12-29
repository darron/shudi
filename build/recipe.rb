class Shudi < FPM::Cookery::Recipe
  name 'shudi'

  version '0.2'
  revision '1'
  description 'shudi'

  homepage 'https://github.com/darron/shudi'
  source "https://github.com/darron/shudi/releases/download/v#{version}/shudi-#{version}-linux-amd64.zip"
  sha256 'b7f52c8ed5952037ebd7fa5f0ffd156251007a2888a18b24b454d9f33ddc6d28'

  maintainer 'Darron <darron@froese.org>'

  license 'Mozilla Public License, version 2.0'

  conflicts 'shudi'
  replaces 'shudi'

  build_depends 'unzip'

  def build
    safesystem "mkdir -p #{builddir}/usr/local/bin/"
    safesystem "cp -f #{builddir}/shudi-#{version}-linux-amd64/shudi-#{version}-linux-amd64 #{builddir}/usr/local/bin/shudi"
  end

  def install
    safesystem "mkdir -p #{destdir}/usr/local/bin/"
    safesystem "cp -f #{builddir}/usr/local/bin/shudi #{destdir}/usr/local/bin/shudi"
    safesystem "chmod 755 #{destdir}/usr/local/bin/shudi"
  end
end
