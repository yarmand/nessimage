#!/usr/bin/env ruby

module Nessimage
  DIRS = {
    fishes: './images/fishes',
    nessies: './images/nessie',
    images: './images'
  }

  class Looper

    attr_reader :nb_images, :output_dir

    def initialize(nb_images,  output_dir)
      @nb_images  = Integer(nb_images)
      @output_dir = output_dir
      puts "nb_images: #{nb_images}"
      generate_images
    end

    def generate_images
      nessies_left = nessies.size
      Range.new(1,nb_images).each do |i|
        if i < nb_images*3/4 || nessies_left <= 0
          generate_nothing_or_fish(i)
        else
          generate_nessie(i, (nessies.size - nessies_left ))
          nessies_left -= 1
        end
      end
    end

    def generate_nothing_or_fish(index)
      if rand(100) < 11
        generate_fish(index)
      else
        generate_nothing(index)
      end
    end

    def generate_nothing(index)
      filename = "#{output_dir}/#{index}_n.png"
      input = "#{DIRS[:images]}/nothing.png"
      generate_ultrasound(input, filename)
    end

    def generate_fish(index)
      filename = "#{output_dir}/#{index}_f.png"
      fish = fishes[rand(fishes.size)]
      resized_fish = "/tmp/tmp_fish.png"
      resize_and_twist(fish, resized_fish)
      input = "/tmp/tmp_input.png"
      put_on_white(resized_fish, input)
      generate_ultrasound(input, filename)
    end

    def resize_and_twist(src, dest)
      system("convert #{src} -resize #{8000 + rand(2000)}@ #{dest}")
    end

    def generate_white(filename)
      system("convert -size 1000x1400 xc:white #{filename}")
    end

    def put_on_white(src, dest)
      system("convert -size 1000x1400 xc:white \
              #{src} -geometry +#{rand(1000)}+#{rand(1400)} \
             -composite #{dest}")
    end

    def generate_nessie(index, nessie_index)
      filename = "#{output_dir}/#{index}_f.png"
      input = nessies[nessie_index]
      generate_ultrasound(input, filename)
    end

    def generate_ultrasound(input, output)
      #system("go run main.go #{input} #{output}")
      system("./main #{input} #{output}")
    end

    def fishes
      @fishes ||= Dir.glob("#{DIRS[:fishes]}/*.png")
    end

    def nessies
      @nessies ||= Dir.glob("#{DIRS[:nessies]}/*.png")
    end
  end

end

NB_IMG_TO_GENERATE = ARGV[0]
OUTPUT_DIR         = ARGV[1]

puts "NB_IMG_TO_GENERATE: #{NB_IMG_TO_GENERATE}"

Nessimage::Looper.new(NB_IMG_TO_GENERATE, OUTPUT_DIR)
